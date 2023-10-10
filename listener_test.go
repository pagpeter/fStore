package fstore

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ExampleStruct struct {
	IP          string `json:"ip"`
	HTTPVersion string `json:"http_version"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	TLS         struct {
		Ciphers    []string `json:"ciphers"`
		Extensions []struct {
			Name                     string   `json:"name"`
			Data                     string   `json:"data,omitempty"`
			SupportedGroups          []string `json:"supported_groups,omitempty"`
			MasterSecretData         string   `json:"master_secret_data,omitempty"`
			ExtendedMasterSecretData string   `json:"extended_master_secret_data,omitempty"`
			SignatureAlgorithms      []string `json:"signature_algorithms,omitempty"`
			StatusRequest            struct {
				CertificateStatusType   string `json:"certificate_status_type"`
				ResponderIDListLength   int    `json:"responder_id_list_length"`
				RequestExtensionsLength int    `json:"request_extensions_length"`
			} `json:"status_request,omitempty"`
			Versions   []string `json:"versions,omitempty"`
			Algorithms []string `json:"algorithms,omitempty"`
			ServerName string   `json:"server_name,omitempty"`
			SharedKeys []struct {
				TLSGREASE0Xfafa     string `json:"TLS_GREASE (0xfafa),omitempty"`
				X25519Kyber76825497 string `json:"X25519Kyber768 (25497),omitempty"`
				X2551929            string `json:"X25519 (29),omitempty"`
			} `json:"shared_keys,omitempty"`
			EllipticCurvesPointFormats []string `json:"elliptic_curves_point_formats,omitempty"`
			Protocols                  []string `json:"protocols,omitempty"`
			PSKKeyExchangeMode         string   `json:"PSK_Key_Exchange_Mode,omitempty"`
		} `json:"extensions"`
		TLSVersionRecord     string `json:"tls_version_record"`
		TLSVersionNegotiated string `json:"tls_version_negotiated"`
		Ja3                  string `json:"ja3"`
		Ja3Hash              string `json:"ja3_hash"`
		Peetprint            string `json:"peetprint"`
		PeetprintHash        string `json:"peetprint_hash"`
		ClientRandom         string `json:"client_random"`
		SessionID            string `json:"session_id"`
	} `json:"tls"`
	HTTP2 struct {
		AkamaiFingerprint     string `json:"akamai_fingerprint"`
		AkamaiFingerprintHash string `json:"akamai_fingerprint_hash"`
		SentFrames            []struct {
			FrameType string   `json:"frame_type"`
			Length    int      `json:"length"`
			Settings  []string `json:"settings,omitempty"`
			Increment int      `json:"increment,omitempty"`
			StreamID  int      `json:"stream_id,omitempty"`
			Headers   []string `json:"headers,omitempty"`
			Flags     []string `json:"flags,omitempty"`
			Priority  struct {
				Weight    int `json:"weight"`
				DependsOn int `json:"depends_on"`
				Exclusive int `json:"exclusive"`
			} `json:"priority,omitempty"`
		} `json:"sent_frames"`
	} `json:"http2"`
}

var fp = []byte(`{
  "ip": "1.1.1.1:59418",
  "http_version": "h2",
  "method": "GET",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
  "tls": {
    "ciphers": [
      "TLS_GREASE (0x8A8A)",
      "TLS_AES_128_GCM_SHA256",
      "TLS_AES_256_GCM_SHA384",
      "TLS_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
      "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
      "TLS_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_RSA_WITH_AES_128_CBC_SHA",
      "TLS_RSA_WITH_AES_256_CBC_SHA"
    ],
    "extensions": [
      {
        "name": "TLS_GREASE (0x9a9a)"
      },
      {
        "name": "signed_certificate_timestamp (18)"
      },
      {
        "name": "session_ticket (35)",
        "data": ""
      },
      {
        "name": "supported_groups (10)",
        "supported_groups": [
          "TLS_GREASE (0xfafa)",
          "X25519Kyber768 (25497)",
          "X25519 (29)",
          "P-256 (23)",
          "P-384 (24)"
        ]
      },
      {
        "name": "extended_master_secret (23)",
        "master_secret_data": "",
        "extended_master_secret_data": ""
      },
      {
        "name": "signature_algorithms (13)",
        "signature_algorithms": [
          "ecdsa_secp256r1_sha256",
          "rsa_pss_rsae_sha256",
          "rsa_pkcs1_sha256",
          "ecdsa_secp384r1_sha384",
          "rsa_pss_rsae_sha384",
          "rsa_pkcs1_sha384",
          "rsa_pss_rsae_sha512",
          "rsa_pkcs1_sha512"
        ]
      },
      {
        "name": "status_request (5)",
        "status_request": {
          "certificate_status_type": "OSCP (1)",
          "responder_id_list_length": 0,
          "request_extensions_length": 0
        }
      },
      {
        "name": "supported_versions (43)",
        "versions": [
          "TLS_GREASE (0xdada)",
          "TLS 1.3",
          "TLS 1.2"
        ]
      },
      {
        "name": "compress_certificate (27)",
        "algorithms": [
          "brotli (2)"
        ]
      },
      {
        "name": "server_name (0)",
        "server_name": "tls.peet.ws"
      },
      {
        "name": "extensionRenegotiationInfo (boringssl) (65281)",
        "data": "00"
      },
      {
        "name": "key_share (51)",
        "shared_keys": [
          {
            "TLS_GREASE (0xfafa)": "00"
          },
          {
            "X25519Kyber768 (25497)": "7b363296460eb1c02706b128953e6d75bc36d71d499de8623379ba831be44f39e31182fb5c724af745adf7012bbc0b8ef9ba6f66124f775b1b5346c2071cdd5a6d7cd6c3bff84291ab87fd735aae6403a3d109b5e18ca7e0b98833259c442a730ba31e6041a0f309fb4b5dcad1b6700a52c6c4421a0616adf43cd0639c4dd8b62f8945c9278cad96c60fb140aa0759e32a1365d06ada077188b33922e038eac0caea79cebc8c921106a33c6c8febc70613a3036b666a365a4bf83b0bdcd142387aa375a49e1352a4eb6656c9c53f5f6812789728067a6e355ccf02aa1b6cb16500024c8a35688c9170b074aeb5e0a70452ca83c94b428b75bec6245ad971d1c854ac2bc041f95293e45e0687ac7c152815c0ab34d79e569b35530736ebd90a6237746426be3a811be448243bdbb0f628c4022514bb04bd56d5113c10486843b0c574c414a176c6c62957032dda84c0dd9ac6c8b9251e2c0caf563dc606b726b135d5926357a3ca93327564276efe8970c648c62439aa2ff03a4f65a1cd8273d73b899967b705585703c47c3b7c80528c1eadb4254222c5bea95020fc114361bb7d65a1f768123e977bc51b82ee4bce2e54bee8b44ab1f929de1171ff950de2048280c8bffdd082e850861bb50fb580ad2b6276e84c529dc34402c761a53811ee206c68d66c7213a88eba6d39a51c3d8a0a05695a422931cc9a27bfbc3052694d29766b587862e8e0cbc126c22f4081cea9700ee989f8ab835d03b858962507f0aaa7734d56c2a867205c3c448e74a95eac7554d0f7c8853044ea73ca369aad310c3a4f06318837515cc5b90f29446dfc05fbac2641f52e6829c22b348c75107ee19a549140b3d7f32d898a1114f58ddc119adf642b31b2b0fdf13831eb4465b4c0246679111761e0aa8f1e135f91acb2eff9b37f6c459297be885b7f8e16c0d4364d33a9c51f9a8aa1b4a30127b97a96a6dc9257902a0d51a32fbb298ca7f0362d54c52959924a758b3110404fd14d28080711a7714a109ccb178dd68424651a75578499b21a7226c203a129281e3300ca9459909a0b5163bc8a39408cda65927365c4d6c7ed8a19ae37b7d7f02e900c02b9e7865cd621d681c4ca438a3d8863eef4cce9126ee95096f4d7833ffc8c895cca79c105d5d65bf26353448b40aeb53da0226fb9f063a66389ae63962f729da684403a67cc096545c7318af872aa713cc7a25260671a360bc8226f287c70e2a174e086fda954152671a996a49ee60d8609a2485618a1d7a9f7b6a2004dc16dd479f79121dbb3c53586aa8e075d8f76be89806b23e28ccc1cb285862bd625b9b0d053afe0b450b9c554589654853236f3bea59645234ccb289025f21603f082640420cfc7551afa653e94f80d45e8b037e8cda8d07074993f4f3068cbb21e464a2745c52dd226610080083d9467f3b831c379a40700147dd340d88b1875e801ea33770c47b6dea5795c3356593209073c6a9e5cc9696b227f013abc12642a91020dfb41dc761dac56368261bf5ddc6dfb96a19d43bce827b7ce41a90157afc7db49062cc034e88ce14b897ccb8c020c867e0550fc498659e84f64bc4772e43f5ee5bfc27bc5d245a66854745fbb460d61b8bc362bc876870bd2ce83009112a49e64739b8c0700b72acbd8a202862722e54051a4d37313b0c439b0f52eb97e77500d83e51ed3adcad746beec3863979b"
          },
          {
            "X25519 (29)": "e43687094584db53755b0a9abc84a7433579aa5598db91ae57950d158139f01b"
          }
        ]
      },
      {
        "name": "ec_point_formats (11)",
        "elliptic_curves_point_formats": [
          "0x00"
        ]
      },
      {
        "name": "application_layer_protocol_negotiation (16)",
        "protocols": [
          "h2",
          "http/1.1"
        ]
      },
      {
        "name": "application_settings (17513)",
        "protocols": [
          "h2"
        ]
      },
      {
        "name": "psk_key_exchange_modes (45)",
        "PSK_Key_Exchange_Mode": "PSK with (EC)DHE key establishment (psk_dhe_ke) (1)"
      },
      {
        "name": "TLS_GREASE (0x7a7a)"
      },
      {
        "name": "pre_shared_key (41)",
        "data": "007700714ba602a57be4dc4461eba9dbc6d54f5405897aeff7339546e52ee92bb71b55719247cac9b198e3af2eaeaa5ffbbc84c34404ea2ae9ba7a0cd839b770d8b537691d74416ca5d868bb61f367beb68c513818ca8d4db3e4c06b9362cc5327caffec6b1fc0323a5ef70d2e4536dbc638d5d8d61a26127900212001247a71878613f7d8ffc455d0b8a356c3a2ddd1c297a86d3f20e39d58d561f4"
      }
    ],
    "tls_version_record": "771",
    "tls_version_negotiated": "772",
    "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,18-35-10-23-13-5-43-27-0-65281-51-11-16-17513-45-41,25497-29-23-24,0",
    "ja3_hash": "9f6db5c1bf0924780372cca95bb7f8ba",
    "peetprint": "GREASE-772-771|2-1.1|GREASE-25497-29-23-24|1027-2052-1025-1283-2053-1281-2054-1537|1|2|GREASE-4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53|0-10-11-13-16-17513-18-23-27-35-41-43-45-5-51-65281-GREASE-GREASE",
    "peetprint_hash": "0cd24525a36f772d1f68f62ea0eb24ce",
    "client_random": "c62cb50847e207bb8c0f391ac9bec7fa47001aa7cc274f9cf1a5764ff1bc1827",
    "session_id": "eb6ddd9d4765afb1c5c780e2087efbce13f1f84b0da7e7b3107230d6216f752d"
  },
  "http2": {
    "akamai_fingerprint": "1:65536,2:0,4:6291456,6:262144|15663105|0|m,a,s,p",
    "akamai_fingerprint_hash": "90224459f8bf70b7d0a8797eb916dbc9",
    "sent_frames": [
      {
        "frame_type": "SETTINGS",
        "length": 24,
        "settings": [
          "HEADER_TABLE_SIZE = 65536",
          "ENABLE_PUSH = 0",
          "INITIAL_WINDOW_SIZE = 6291456",
          "MAX_HEADER_LIST_SIZE = 262144"
        ]
      },
      {
        "frame_type": "WINDOW_UPDATE",
        "length": 4,
        "increment": 15663105
      },
      {
        "frame_type": "HEADERS",
        "stream_id": 1,
        "length": 568,
        "headers": [
          ":method: GET",
          ":authority: tls.peet.ws",
          ":scheme: https",
          ":path: /api/all",
          "sec-ch-ua: \\\"Google Chrome\\\";v=\\\"117\\\", \\\"Not;A=Brand\\\";v=\\\"8\\\", \\\"Chromium\\\";v=\\\"117\\",
          "sec-ch-ua-mobile: ?0",
          "sec-ch-ua-platform: \\\"macOS\\",
          "upgrade-insecure-requests: 1",
          "dnt: 1",
          "user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
          "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
          "sec-fetch-site: same-origin",
          "sec-fetch-mode: navigate",
          "sec-fetch-user: ?1",
          "sec-fetch-dest: document",
          "referer: https://tls.peet.ws/",
          "accept-encoding: gzip, deflate, br",
          "accept-language: en-DE,en-US;q=0.9,en;q=0.8",
          "cookie: cf_clearance=MUoJKskype.U8euK.Lj3E9nLBa3Vfyc9eQNl7vpQKsk-1696544599-0-1-957868af.5df2091a.69ec06d1-160.2.1696544599"
        ],
        "flags": [
          "EndStream (0x1)",
          "EndHeaders (0x4)",
          "Priority (0x20)"
        ],
        "priority": {
          "weight": 256,
          "depends_on": 0,
          "exclusive": 1
        }
      }
    ]
  }
}`)

// func TestString(m *testing.T) {
// 	f := Listener()
// 	f.EnableDebug()
// 	f.Store(string(fp))
// }

func TestStruct(m *testing.T) {
	var data = &ExampleStruct{}

	json.Unmarshal(fp, data)

	f := Listener()
	f.EnableDebug()
	f.DontHash = []string{
		"client_random", "session_id", "pre_shared_key", "X25519Kyber768",
	}
	f.Threshhold = 5
	f.UseKeyCompression = false

	res, _ := f.Store(data)
	smolData, _ := json.Marshal(res)
	valLookup, _ := json.Marshal(f.database.hashValues)
	keyLookup, _ := json.Marshal(f.database.hashKeys)
	fmt.Println("Data:", string(smolData))
	fmt.Print("\n==\n\n")
	fmt.Println("val lookup:", string(valLookup))
	fmt.Print("\n==\n\n")
	fmt.Println("key lookup:", string(keyLookup))
}
