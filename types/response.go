package types

type AuthSignInfo interface{}

type CertifiedTransaction struct {
	TransactionDigest string        `json:"transactionDigest"`
	TxSignature       string        `json:"txSignature"`
	AuthSignInfo      *AuthSignInfo `json:"authSignInfo"`

	Data *SenderSignedData `json:"data"`
}

type GasCostSummary struct {
	ComputationCost uint64 `json:"computationCost"`
	StorageCost     uint64 `json:"storageCost"`
	StorageRebate   uint64 `json:"storageRebate"`
}

const (
	TransactionStatusSuccess = "success"
	TransactionStatusFailure = "failure"
)

type TransactionStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type TransactionEffects struct {
	Status TransactionStatus `json:"status"`

	TransactionDigest string          `json:"transactionDigest"`
	GasUsed           *GasCostSummary `json:"gasUsed"`
	GasObject         *OwnedObjectRef `json:"gasObject"`
	Events            []Event         `json:"events,omitempty"`
	Dependencies      []string        `json:"dependencies,omitempty"`

	// SharedObjects []ObjectRef      `json:"sharedObjects"`
	Created   []OwnedObjectRef `json:"created,omitempty"`
	Mutated   []OwnedObjectRef `json:"mutated,omitempty"`
	Unwrapped []OwnedObjectRef `json:"unwrapped,omitempty"`
	Deleted   []ObjectRef      `json:"deleted,omitempty"`
	Wrapped   []ObjectRef      `json:"wrapped,omitempty"`
}

func (te *TransactionEffects) GasFee() uint64 {
	return te.GasUsed.StorageCost - te.GasUsed.StorageRebate + te.GasUsed.ComputationCost
}

type ParsedTransactionResponse interface{}

// TxnRequestTypeImmediateReturn       ExecuteTransactionRequestType = "ImmediateReturn"
// TxnRequestTypeWaitForTxCert         ExecuteTransactionRequestType = "WaitForTxCert"
// TxnRequestTypeWaitForEffectsCert    ExecuteTransactionRequestType = "WaitForEffectsCert"

//type TransactionResponse struct {
//	Certificate *CertifiedTransaction     `json:"certificate"`
//	Effects     *TransactionEffects       `json:"effects"`
//	ParsedData  ParsedTransactionResponse `json:"parsed_data,omitempty"`
//	TimestampMs uint64                    `json:"timestamp_ms,omitempty"`
//}
type TransactionResponse struct {
	Certificate struct {
		TransactionDigest string `json:"transactionDigest"`
		Data              struct {
			Transactions []struct {
				Call struct {
					Package   string `json:"package"`
					Module    string `json:"module"`
					Function  string `json:"function"`
					Arguments []any  `json:"arguments"`
				} `json:"Call"`
			} `json:"transactions"`
			Sender  string `json:"sender"`
			GasData struct {
				Payment struct {
					ObjectID string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"payment"`
				Owner  string `json:"owner"`
				Price  int    `json:"price"`
				Budget int    `json:"budget"`
			} `json:"gasData"`
		} `json:"data"`
		TxSignatures []string `json:"txSignatures"`
		AuthSignInfo struct {
			Epoch      int    `json:"epoch"`
			Signature  string `json:"signature"`
			SignersMap []int  `json:"signers_map"`
		} `json:"authSignInfo"`
	} `json:"certificate"`
	Effects struct {
		Status struct {
			Status string `json:"status"`
		} `json:"status"`
		ExecutedEpoch int `json:"executedEpoch"`
		GasUsed       struct {
			ComputationCost int `json:"computationCost"`
			StorageCost     int `json:"storageCost"`
			StorageRebate   int `json:"storageRebate"`
		} `json:"gasUsed"`
		SharedObjects []struct {
			ObjectID string `json:"objectId"`
			Version  int    `json:"version"`
			Digest   string `json:"digest"`
		} `json:"sharedObjects"`
		TransactionDigest string `json:"transactionDigest"`
		Mutated           []struct {
			Owner struct {
				Shared struct {
					InitialSharedVersion int `json:"initial_shared_version"`
				} `json:"Shared"`
			} `json:"owner,omitempty"`
			Reference struct {
				ObjectID string `json:"objectId"`
				Version  int    `json:"version"`
				Digest   string `json:"digest"`
			} `json:"reference"`
			Owner0 struct {
				AddressOwner string `json:"AddressOwner"`
			} `json:"owner,omitempty"`
		} `json:"mutated"`
		GasObject struct {
			Owner struct {
				AddressOwner string `json:"AddressOwner"`
			} `json:"owner"`
			Reference struct {
				ObjectID string `json:"objectId"`
				Version  int    `json:"version"`
				Digest   string `json:"digest"`
			} `json:"reference"`
		} `json:"gasObject"`
		Events []struct {
			CoinBalanceChange *struct {
				PackageID         string `json:"packageId"`
				TransactionModule string `json:"transactionModule"`
				Sender            string `json:"sender"`
				ChangeType        string `json:"changeType"`
				Owner             struct {
					AddressOwner string `json:"AddressOwner"`
				} `json:"owner"`
				CoinType     string `json:"coinType"`
				CoinObjectID string `json:"coinObjectId"`
				Version      int    `json:"version"`
				Amount       int    `json:"amount"`
			} `json:"coinBalanceChange,omitempty"`
			MutateObject *struct {
				PackageID         string `json:"packageId"`
				TransactionModule string `json:"transactionModule"`
				Sender            string `json:"sender"`
				ObjectType        string `json:"objectType"`
				ObjectID          string `json:"objectId"`
				Version           int    `json:"version"`
			} `json:"mutateObject,omitempty"`
			MoveEvent *struct {
				PackageID         string `json:"packageId"`
				TransactionModule string `json:"transactionModule"`
				Sender            string `json:"sender"`
				Type              string `json:"type"`
				Fields            map[string]interface{} `json:"fields"`
				Bcs string `json:"bcs"`
			} `json:"moveEvent,omitempty"`
		} `json:"events"`
		Dependencies []string `json:"dependencies"`
	} `json:"effects"`
	TimestampMs int64 `json:"timestamp_ms"`
	Checkpoint  int   `json:"checkpoint"`
	ParsedData  any   `json:"parsed_data"`
}

type ExecuteTransactionEffects struct {
	TransactionEffectsDigest string `json:"transactionEffectsDigest"`

	Effects      TransactionEffects `json:"effects"`
	AuthSignInfo *AuthSignInfo      `json:"authSignInfo"`
}

func (r *ExecuteTransactionResponse) TransactionDigest() string {
	return r.Certificate.TransactionDigest
}

type ExecuteTransactionResponse  struct {
	Certificate struct {
		TransactionDigest string `json:"transactionDigest"`
		Data              struct {
			Transactions []struct {
				Publish struct {
					Disassembled map[string]interface{} `json:"disassembled"`
				} `json:"Publish"`
			} `json:"transactions"`
			Sender  string `json:"sender"`
			GasData struct {
				Payment struct {
					ObjectID string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"payment"`
				Owner  string `json:"owner"`
				Price  int    `json:"price"`
				Budget int    `json:"budget"`
			} `json:"gasData"`
		} `json:"data"`
		TxSignatures []string `json:"txSignatures"`
		AuthSignInfo struct {
			Epoch      int    `json:"epoch"`
			Signature  string `json:"signature"`
			SignersMap []int  `json:"signers_map"`
		} `json:"authSignInfo"`
	} `json:"certificate"`
	Effects struct {
		TransactionEffectsDigest string `json:"transactionEffectsDigest"`
		Effects                  struct {
			Status struct {
				Status string `json:"status"`
			} `json:"status"`
			ExecutedEpoch int `json:"executedEpoch"`
			GasUsed       struct {
				ComputationCost int `json:"computationCost"`
				StorageCost     int `json:"storageCost"`
				StorageRebate   int `json:"storageRebate"`
			} `json:"gasUsed"`
			TransactionDigest string `json:"transactionDigest"`
			Created           []struct {
				Owner     interface{} `json:"owner"`
				Reference struct {
					ObjectID string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"reference"`
			} `json:"created"`
			Mutated []struct {
				Owner struct {
					AddressOwner string `json:"AddressOwner"`
				} `json:"owner"`
				Reference struct {
					ObjectID string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"reference"`
			} `json:"mutated"`
			GasObject struct {
				Owner struct {
					AddressOwner string `json:"AddressOwner"`
				} `json:"owner"`
				Reference struct {
					ObjectID string `json:"objectId"`
					Version  int    `json:"version"`
					Digest   string `json:"digest"`
				} `json:"reference"`
			} `json:"gasObject"`
			Events []struct {
				CoinBalanceChange *struct {
					PackageID         string `json:"packageId"`
					TransactionModule string `json:"transactionModule"`
					Sender            string `json:"sender"`
					ChangeType        string `json:"changeType"`
					Owner             struct {
						AddressOwner string `json:"AddressOwner"`
					} `json:"owner"`
					CoinType     string `json:"coinType"`
					CoinObjectID string `json:"coinObjectId"`
					Version      int    `json:"version"`
					Amount       int    `json:"amount"`
				} `json:"coinBalanceChange,omitempty"`
				NewObject *struct {
					PackageID         string `json:"packageId"`
					TransactionModule string `json:"transactionModule"`
					Sender            string `json:"sender"`
					Recipient         interface{} `json:"recipient"`
					ObjectType string `json:"objectType"`
					ObjectID   string `json:"objectId"`
					Version    int    `json:"version"`
				} `json:"newObject,omitempty"`
				Publish *struct {
					Sender    string `json:"sender"`
					PackageID string `json:"packageId"`
					Version   int    `json:"version"`
					Digest    string `json:"digest"`
				} `json:"publish,omitempty"`
			} `json:"events"`
			Dependencies []string `json:"dependencies"`
		} `json:"effects"`
		FinalityInfo struct {
			Certified struct {
				Epoch      int    `json:"epoch"`
				Signature  string `json:"signature"`
				SignersMap []int  `json:"signers_map"`
			} `json:"certified"`
		} `json:"finalityInfo"`
	} `json:"effects"`
	ConfirmedLocalExecution bool `json:"confirmed_local_execution"`

}



type SuiCoinMetadata struct {
	Decimals    uint8    `json:"decimals"`
	Description string   `json:"description"`
	IconUrl     string   `json:"iconUrl,omitempty"`
	Id          ObjectId `json:"id"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
}

type SuiCoinBalance struct {
	CoinType        string `json:"coinType"`
	CoinObjectCount int64  `json:"coinObjectCount"`
	TotalBalance    int64  `json:"totalBalance"`
}

type DevInspectResults struct {
	Effects TransactionEffects `json:"effects"`
	Results DevInspectResult   `json:"results"`
}

type DevInspectResult struct {
	Err string `json:"Err,omitempty"`
	Ok  any    `json:"Ok,omitempty"` //Result_of_Array_of_Tuple_of_uint_and_SuiExecutionResult_or_String
}

type CoinPage struct {
	Data       []CoinObject `json:"data"`
	NextCursor *ObjectId    `json:"nextCursor"`
}

type CoinObject struct {
	CoinType     string   `json:"coinType"`
	CoinObjectId ObjectId `json:"coinObjectId"`
	Version      uint64   `json:"version"`
	Digest       Digest   `json:"digest"`
	Balance      int64    `json:"balance"`
}

type Supply struct {
	Value uint64 `json:"value"`
}

type TransactionsPage struct {
	Data       []string `json:"data"`
	NextCursor string   `json:"nextCursor"`
}

type EventPage struct {
	Data       []Event `json:"data"`
	NextCursor EventID `json:"nextCursor"`
}

type EventID struct {
	TxDigest Digest `json:"txDigest"`
	EventSeq int64  `json:"eventSeq"`
}
