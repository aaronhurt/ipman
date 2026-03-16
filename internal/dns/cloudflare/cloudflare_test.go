package cloudflare

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/leprechau/ipman/internal/dns"
)

func TestGetClearsStaleRecordIDAndUpsertCreatesRecord(t *testing.T) {
	t.Parallel()

	var (
		mu      sync.Mutex
		methods []string
		paths   []string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		methods = append(methods, r.Method)
		paths = append(paths, r.URL.Path)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/zones/test-zone/dns_records":
			_ = json.NewEncoder(w).Encode(DNSResponse{
				Success: true,
				Result:  []DNSRecord{},
				ResultInfo: ResponseInfo{
					Count: 0,
				},
			})
		case r.Method == http.MethodPost && r.URL.Path == "/zones/test-zone/dns_records":
			var record DNSRecord
			if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
				t.Fatalf("decode request body: %v", err)
			}

			if record.Content != "1.2.3.4" {
				t.Fatalf("POST content = %q, want %q", record.Content, "1.2.3.4")
			}

			_ = json.NewEncoder(w).Encode(DNSUpdateResponse{
				Success: true,
				Result: DNSRecord{
					Name: record.Name,
				},
			})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	cfg := DefaultConfig()
	cfg.client.SetBaseURL(server.URL)
	cfg.recordID = "stale-id"

	got, err := cfg.Get("test-zone", "www.example.com", dns.A)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if got != "" {
		t.Fatalf("Get() = %q, want empty result for missing record", got)
	}

	if cfg.recordID != "" {
		t.Fatalf("recordID after Get() = %q, want cleared stale value", cfg.recordID)
	}

	name, err := cfg.Upsert("test-zone", "www.example.com", "1.2.3.4", dns.A)
	if err != nil {
		t.Fatalf("Upsert() error = %v", err)
	}

	if name != "www.example.com" {
		t.Fatalf("Upsert() = %q, want %q", name, "www.example.com")
	}

	mu.Lock()
	defer mu.Unlock()

	if len(methods) != 2 {
		t.Fatalf("request count = %d, want 2", len(methods))
	}

	if methods[0] != http.MethodGet || methods[1] != http.MethodPost {
		t.Fatalf("methods = %v, want [GET POST]", methods)
	}

	if strings.Contains(paths[1], "stale-id") {
		t.Fatalf("POST path = %q, stale record id leaked into create path", paths[1])
	}
}

func TestUpsertUpdatesExistingRecord(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Fatalf("method = %s, want PATCH", r.Method)
		}

		if r.URL.Path != "/zones/test-zone/dns_records/record-123" {
			t.Fatalf("path = %q, want %q", r.URL.Path, "/zones/test-zone/dns_records/record-123")
		}

		var record DNSRecord
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			t.Fatalf("decode request body: %v", err)
		}

		if record.ID != "record-123" {
			t.Fatalf("request body id = %q, want %q", record.ID, "record-123")
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DNSUpdateResponse{
			Success: true,
			Result: DNSRecord{
				Name: record.Name,
			},
		})
	}))
	defer server.Close()

	cfg := DefaultConfig()
	cfg.client.SetBaseURL(server.URL)
	cfg.recordID = "record-123"

	name, err := cfg.Upsert("test-zone", "www.example.com", "1.2.3.4", dns.A)
	if err != nil {
		t.Fatalf("Upsert() error = %v", err)
	}

	if name != "www.example.com" {
		t.Fatalf("Upsert() = %q, want %q", name, "www.example.com")
	}
}
