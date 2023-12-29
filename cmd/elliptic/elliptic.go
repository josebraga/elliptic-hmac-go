package main

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
  "fmt"
  "log"
  "strconv"
  "strings"
)

// Generate a signature for use when signing a request to the API
// - secret:          your secret supplied by Elliptic - a base64 encoded string
// - time_of_request: current time, in milliseconds, since 1 Jan 1970 00:00:00 UTC
// - http_method:     must be uppercase
// - http_path:       API endpoint including query string
// - payload:         string encoded JSON object or '{}' if there is no request body
func get_signature(secret string, time_of_request int64, http_method string, http_path string, payload string) string {
  // create a SHA256 HMAC using the supplied secret, decoded from base64
  ds, err := base64.StdEncoding.DecodeString(secret)
  if err != nil {
    log.Fatal("error:", err)
  }
  h := hmac.New(sha256.New, []byte(ds))

  // concatenate the request text to be signed
  request_text := strconv.FormatInt(time_of_request, 10) + http_method + strings.ToLower(http_path) + payload

  // update the HMAC with the text to be signed
  h.Write([]byte(request_text))

  // output the signature as a base64 encoded string
  return base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

func main() {
  secret := "894f142d667e8cdaca6822ac173937af" // Supplied by Elliptic
  // Disclaimer: this secret is just an example
  time_of_request_in_ms := int64(1478692862000) // For real world use time.Now().UnixMilli()

  example_payload := `[{"customer_reference":"123456","subject":{"asset":"BTC","hash":"accf5c09cc027339a3beb2e28104ce9f406ecbbd29775b4a1a17ba213f1e035e","output_address":"15Hm2UEPaEuiAmgyNgd5mF3wugqLsYs3Wn","output_type":"address","type":"transaction"},"type":"source_of_funds"}]`

  // Example One: POST with payload - you only need to run stringify json when passing a request body
  fmt.Println(get_signature(secret, time_of_request_in_ms, "POST", "/v2/analyses", example_payload))
  // 65mQHB2o95lL3I+N/bZYwDC9p2YvNwsVDnXr8u72hUk=

  // Example Two: GET with empty payload - do not run stringify with no request body, pass an empty object as string
  fmt.Println(get_signature(secret, time_of_request_in_ms, "GET", "/v2/customers", `{}`))
  // cN9fRUqeT7UnwwpkBZaNmnwxKAPHkhytdXelfUVvxMI=
}

