package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/PolicyCRUD/CreateOrUpdateBackupPolicy.json
func ExampleBackupPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "000pikumar", "PrivatePreviewVault", "OSSDBPolicy", armdataprotection.BaseBackupPolicyResource{
		Properties: &armdataprotection.BackupPolicy{
			DatasourceTypes: []*string{
				to.Ptr("OssDB")},
			ObjectType: to.Ptr("BackupPolicy"),
			PolicyRules: []armdataprotection.BasePolicyRuleClassification{
				&armdataprotection.AzureBackupRule{
					Name:       to.Ptr("BackupWeekly"),
					ObjectType: to.Ptr("AzureBackupRule"),
					BackupParameters: &armdataprotection.AzureBackupParams{
						ObjectType: to.Ptr("AzureBackupParams"),
						BackupType: to.Ptr("Full"),
					},
					DataStore: &armdataprotection.DataStoreInfoBase{
						DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
						ObjectType:    to.Ptr("DataStoreInfoBase"),
					},
					Trigger: &armdataprotection.ScheduleBasedTriggerContext{
						ObjectType: to.Ptr("ScheduleBasedTriggerContext"),
						Schedule: &armdataprotection.BackupSchedule{
							RepeatingTimeIntervals: []*string{
								to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
						},
						TaggingCriteria: []*armdataprotection.TaggingCriteria{
							{
								IsDefault: to.Ptr(true),
								TagInfo: &armdataprotection.RetentionTag{
									TagName: to.Ptr("Default"),
								},
								TaggingPriority: to.Ptr[int64](99),
							},
							{
								Criteria: []armdataprotection.BackupCriteriaClassification{
									&armdataprotection.ScheduleBasedBackupCriteria{
										ObjectType: to.Ptr("ScheduleBasedBackupCriteria"),
										DaysOfTheWeek: []*armdataprotection.DayOfWeek{
											to.Ptr(armdataprotection.DayOfWeekSunday)},
										ScheduleTimes: []*time.Time{
											to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t }())},
									}},
								IsDefault: to.Ptr(false),
								TagInfo: &armdataprotection.RetentionTag{
									TagName: to.Ptr("Weekly"),
								},
								TaggingPriority: to.Ptr[int64](20),
							}},
					},
				},
				&armdataprotection.AzureRetentionRule{
					Name:       to.Ptr("Default"),
					ObjectType: to.Ptr("AzureRetentionRule"),
					IsDefault:  to.Ptr(true),
					Lifecycles: []*armdataprotection.SourceLifeCycle{
						{
							DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
								Duration:   to.Ptr("P1W"),
								ObjectType: to.Ptr("AbsoluteDeleteOption"),
							},
							SourceDataStore: &armdataprotection.DataStoreInfoBase{
								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
								ObjectType:    to.Ptr("DataStoreInfoBase"),
							},
						}},
				},
				&armdataprotection.AzureRetentionRule{
					Name:       to.Ptr("Weekly"),
					ObjectType: to.Ptr("AzureRetentionRule"),
					IsDefault:  to.Ptr(false),
					Lifecycles: []*armdataprotection.SourceLifeCycle{
						{
							DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
								Duration:   to.Ptr("P12W"),
								ObjectType: to.Ptr("AbsoluteDeleteOption"),
							},
							SourceDataStore: &armdataprotection.DataStoreInfoBase{
								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
								ObjectType:    to.Ptr("DataStoreInfoBase"),
							},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
