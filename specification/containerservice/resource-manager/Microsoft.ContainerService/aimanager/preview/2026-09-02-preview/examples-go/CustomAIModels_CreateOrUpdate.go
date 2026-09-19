package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-09-02-preview/CustomAIModels_CreateOrUpdate.json
func ExampleCustomAIModelsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomAIModelsClient().BeginCreateOrUpdate(ctx, "rg1", "aimanager1", "custom-model1", armcontainerserviceaimanager.CustomAIModel{
		Properties: &armcontainerserviceaimanager.CustomAIModelProperties{
			ModelID: to.Ptr("Qwen/Qwen36-27B-private"),
			BaseModel: &armcontainerserviceaimanager.BaseModelReference{
				ID:                   to.Ptr("Qwen/Qwen3.6-27B"),
				TotalWeightSizeBytes: to.Ptr[int64](64700000000),
				Config: map[string]any{
					"architectures": []any{
						"Qwen3_5ForConditionalGeneration",
					},
					"image_token_id":      248056,
					"language_model_only": false,
					"model_type":          "qwen3_5",
					"text_config": map[string]any{
						"attention_bias":          false,
						"attention_dropout":       0,
						"attn_output_gate":        true,
						"bos_token_id":            248044,
						"dtype":                   "bfloat16",
						"eos_token_id":            248044,
						"full_attention_interval": 4,
						"head_dim":                256,
						"hidden_act":              "silu",
						"hidden_size":             5120,
						"initializer_range":       0.02,
						"intermediate_size":       17408,
						"layer_types": []any{
							"linear_attention",
							"linear_attention",
							"...",
							"full_attention",
						},
						"linear_conv_kernel_dim":       4,
						"linear_key_head_dim":          128,
						"linear_num_key_heads":         16,
						"linear_num_value_heads":       48,
						"linear_value_head_dim":        128,
						"mamba_ssm_dtype":              "float32",
						"max_position_embeddings":      262144,
						"model_type":                   "qwen3_5_text",
						"mtp_num_hidden_layers":        1,
						"mtp_use_dedicated_embeddings": false,
						"num_attention_heads":          24,
						"num_hidden_layers":            64,
						"num_key_value_heads":          4,
						"output_gate_type":             "swish",
						"pad_token_id":                 nil,
						"partial_rotary_factor":        0.25,
						"rms_norm_eps":                 0.000001,
						"rope_parameters": map[string]any{
							"mrope_interleaved": true,
							"mrope_section": []any{
								11,
								11,
								10,
							},
							"partial_rotary_factor": 0.25,
							"rope_theta":            10000000,
							"rope_type":             "default",
						},
						"tie_word_embeddings": false,
						"use_cache":           true,
						"vocab_size":          248320,
					},
					"tie_word_embeddings":  false,
					"transformers_version": "4.57.1",
					"video_token_id":       248057,
					"vision_config": map[string]any{
						"deepstack_visual_indexes": []any{},
						"depth":                    27,
						"hidden_act":               "gelu_pytorch_tanh",
						"hidden_size":              1152,
						"in_channels":              3,
						"initializer_range":        0.02,
						"intermediate_size":        4304,
						"model_type":               "qwen3_5",
						"num_heads":                16,
						"num_position_embeddings":  2304,
						"out_hidden_size":          5120,
						"patch_size":               16,
						"spatial_merge_size":       2,
						"temporal_patch_size":      2,
					},
					"vision_end_token_id":   248054,
					"vision_start_token_id": 248053,
				},
			},
			ModelSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/foundry-private-source"),
			Description:           to.Ptr("Custom Llama 2 7B model for our organization"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerserviceaimanager.CustomAIModelsClientCreateOrUpdateResponse{
	// 	CustomAIModel: armcontainerserviceaimanager.CustomAIModel{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1/customAIModels/custom-model1"),
	// 		Name: to.Ptr("custom-model1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/aiManagers/customAIModels"),
	// 		SystemData: &armcontainerserviceaimanager.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 		Properties: &armcontainerserviceaimanager.CustomAIModelProperties{
	// 			ProvisioningState: to.Ptr(armcontainerserviceaimanager.CustomAIModelProvisioningStateSucceeded),
	// 			ModelID: to.Ptr("Qwen/Qwen36-27B-private"),
	// 			BaseModel: &armcontainerserviceaimanager.BaseModelReference{
	// 				ID: to.Ptr("Qwen/Qwen3.6-27B"),
	// 				TotalWeightSizeBytes: to.Ptr[int64](64700000000),
	// 				Config: map[string]any{
	// 					"architectures": []any{
	// 						"Qwen3_5ForConditionalGeneration",
	// 					},
	// 					"image_token_id": 248056,
	// 					"language_model_only": false,
	// 					"model_type": "qwen3_5",
	// 					"text_config": map[string]any{
	// 						"attention_bias": false,
	// 						"attention_dropout": 0,
	// 						"attn_output_gate": true,
	// 						"bos_token_id": 248044,
	// 						"dtype": "bfloat16",
	// 						"eos_token_id": 248044,
	// 						"full_attention_interval": 4,
	// 						"head_dim": 256,
	// 						"hidden_act": "silu",
	// 						"hidden_size": 5120,
	// 						"initializer_range": 0.02,
	// 						"intermediate_size": 17408,
	// 						"layer_types": []any{
	// 							"linear_attention",
	// 							"linear_attention",
	// 							"...",
	// 							"full_attention",
	// 						},
	// 						"linear_conv_kernel_dim": 4,
	// 						"linear_key_head_dim": 128,
	// 						"linear_num_key_heads": 16,
	// 						"linear_num_value_heads": 48,
	// 						"linear_value_head_dim": 128,
	// 						"mamba_ssm_dtype": "float32",
	// 						"max_position_embeddings": 262144,
	// 						"model_type": "qwen3_5_text",
	// 						"mtp_num_hidden_layers": 1,
	// 						"mtp_use_dedicated_embeddings": false,
	// 						"num_attention_heads": 24,
	// 						"num_hidden_layers": 64,
	// 						"num_key_value_heads": 4,
	// 						"output_gate_type": "swish",
	// 						"pad_token_id": nil,
	// 						"partial_rotary_factor": 0.25,
	// 						"rms_norm_eps": 0.000001,
	// 						"rope_parameters": map[string]any{
	// 							"mrope_interleaved": true,
	// 							"mrope_section": []any{
	// 								11,
	// 								11,
	// 								10,
	// 							},
	// 							"partial_rotary_factor": 0.25,
	// 							"rope_theta": 10000000,
	// 							"rope_type": "default",
	// 						},
	// 						"tie_word_embeddings": false,
	// 						"use_cache": true,
	// 						"vocab_size": 248320,
	// 					},
	// 					"tie_word_embeddings": false,
	// 					"transformers_version": "4.57.1",
	// 					"video_token_id": 248057,
	// 					"vision_config": map[string]any{
	// 						"deepstack_visual_indexes": []any{
	// 						},
	// 						"depth": 27,
	// 						"hidden_act": "gelu_pytorch_tanh",
	// 						"hidden_size": 1152,
	// 						"in_channels": 3,
	// 						"initializer_range": 0.02,
	// 						"intermediate_size": 4304,
	// 						"model_type": "qwen3_5",
	// 						"num_heads": 16,
	// 						"num_position_embeddings": 2304,
	// 						"out_hidden_size": 5120,
	// 						"patch_size": 16,
	// 						"spatial_merge_size": 2,
	// 						"temporal_patch_size": 2,
	// 					},
	// 					"vision_end_token_id": 248054,
	// 					"vision_start_token_id": 248053,
	// 				},
	// 			},
	// 			ModelSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/foundry-private-source"),
	// 			Description: to.Ptr("Custom Llama 2 7B model for our organization"),
	// 			Spec: &armcontainerserviceaimanager.CustomAIModelSpec{
	// 				License: to.Ptr("llama2"),
	// 				IsRestricted: to.Ptr(false),
	// 				MaxContextLength: to.Ptr[int32](4096),
	// 			},
	// 		},
	// 	},
	// }
}
