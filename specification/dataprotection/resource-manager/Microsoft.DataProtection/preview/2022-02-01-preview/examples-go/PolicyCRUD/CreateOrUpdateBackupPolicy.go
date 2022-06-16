package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/PolicyCRUD/CreateOrUpdateBackupPolicy.json
func ExampleBackupPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-policy-name>",
		armdataprotection.BaseBackupPolicyResource{
			Properties: &armdataprotection.BackupPolicy{
				DatasourceTypes: []*string{
					to.Ptr("OssDB")},
				ObjectType: to.Ptr("<object-type>"),
				PolicyRules: []armdataprotection.BasePolicyRuleClassification{
					&armdataprotection.AzureBackupRule{
						Name:       to.Ptr("<name>"),
						ObjectType: to.Ptr("<object-type>"),
						BackupParameters: &armdataprotection.AzureBackupParams{
							ObjectType: to.Ptr("<object-type>"),
							BackupType: to.Ptr("<backup-type>"),
						},
						DataStore: &armdataprotection.DataStoreInfoBase{
							DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
							ObjectType:    to.Ptr("<object-type>"),
						},
						Trigger: &armdataprotection.ScheduleBasedTriggerContext{
							ObjectType: to.Ptr("<object-type>"),
							Schedule: &armdataprotection.BackupSchedule{
								RepeatingTimeIntervals: []*string{
									to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
							},
							TaggingCriteria: []*armdataprotection.TaggingCriteria{
								{
									IsDefault: to.Ptr(true),
									TagInfo: &armdataprotection.RetentionTag{
										TagName: to.Ptr("<tag-name>"),
									},
									TaggingPriority: to.Ptr[int64](99),
								},
								{
									Criteria: []armdataprotection.BackupCriteriaClassification{
										&armdataprotection.ScheduleBasedBackupCriteria{
											ObjectType: to.Ptr("<object-type>"),
											DaysOfTheWeek: []*armdataprotection.DayOfWeek{
												to.Ptr(armdataprotection.DayOfWeekSunday)},
											ScheduleTimes: []*time.Time{
												to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t }())},
										}},
									IsDefault: to.Ptr(false),
									TagInfo: &armdataprotection.RetentionTag{
										TagName: to.Ptr("<tag-name>"),
									},
									TaggingPriority: to.Ptr[int64](20),
								}},
						},
					},
					&armdataprotection.AzureRetentionRule{
						Name:       to.Ptr("<name>"),
						ObjectType: to.Ptr("<object-type>"),
						IsDefault:  to.Ptr(true),
						Lifecycles: []*armdataprotection.SourceLifeCycle{
							{
								DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
									Duration:   to.Ptr("<duration>"),
									ObjectType: to.Ptr("<object-type>"),
								},
								SourceDataStore: &armdataprotection.DataStoreInfoBase{
									DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
									ObjectType:    to.Ptr("<object-type>"),
								},
							}},
					},
					&armdataprotection.AzureRetentionRule{
						Name:       to.Ptr("<name>"),
						ObjectType: to.Ptr("<object-type>"),
						IsDefault:  to.Ptr(false),
						Lifecycles: []*armdataprotection.SourceLifeCycle{
							{
								DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
									Duration:   to.Ptr("<duration>"),
									ObjectType: to.Ptr("<object-type>"),
								},
								SourceDataStore: &armdataprotection.DataStoreInfoBase{
									DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
									ObjectType:    to.Ptr("<object-type>"),
								},
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
