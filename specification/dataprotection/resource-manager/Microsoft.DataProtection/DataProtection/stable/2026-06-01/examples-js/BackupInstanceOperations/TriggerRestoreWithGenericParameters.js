const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to triggers restore for a BackupInstance
 *
 * @summary triggers restore for a BackupInstance
 * x-ms-original-file: 2026-06-01/BackupInstanceOperations/TriggerRestoreWithGenericParameters.json
 */
async function triggerRestoreWithGenericRestoreDatasourceCriteria() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.triggerRestore(
    "000pikumar",
    "PrivatePreviewVault1",
    "testInstance1",
    {
      objectType: "AzureBackupRecoveryPointBasedRestoreRequest",
      recoveryPointId: "hardcodedRP",
      sourceDataStoreType: "OperationalStore",
      restoreTargetInfo: {
        restoreLocation: "southeastasia",
        recoveryOption: "FailIfExists",
        objectType: "ItemLevelRestoreTargetInfo",
        datasourceInfo: {
          resourceID:
            "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/target-esan-volgroup",
          resourceUri: "SampleresourceUri123",
          datasourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceName: "target-esan-volgroup",
          resourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceLocation: "eastus2euap",
          objectType: "Datasource",
        },
        datasourceSetInfo: {
          resourceID:
            "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc",
          datasourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceType: "Microsoft.ElasticSan/elasticSans",
          resourceLocation: "eastus2euap",
          objectType: "DatasourceSet",
        },
        restoreCriteria: [
          {
            objectType: "GenericRestoreDatasourceCriteria",
            resourceSelectors: {
              objectType: "resourceListSelectionCriteria",
              resourceIdentifiers: ["source-vol1", "source-vol2", "source-vol3"],
              resourceNameOverrides: { "source-vol1": "target-vol1", "source-vol2": "target-vol2" },
            },
          },
        ],
      },
    },
  );
  console.log(result);
}
