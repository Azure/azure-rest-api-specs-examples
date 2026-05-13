const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a backup instance in a backup vault
 *
 * @summary create or update a backup instance in a backup vault
 * x-ms-original-file: 2026-03-01/BackupInstanceOperations/PutBackupInstance_BlobBackupAutoProtection.json
 */
async function createBackupInstanceWithBlobBackupAutoProtection() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "54707983-993e-43de-8d94-074451394eda";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.createOrUpdate(
    "blobrg",
    "blobvault",
    "blobstorageaccount-blobstorageaccount-2a76f8a-c176-4f7d-819e-95157e2b0071",
    {
      properties: {
        dataSourceInfo: {
          datasourceType: "Microsoft.Storage/storageAccounts/blobServices",
          objectType: "Datasource",
          resourceID:
            "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount",
          resourceLocation: "centraluseuap",
          resourceName: "blobstorageaccount",
          resourceType: "microsoft.storage/storageAccounts",
          resourceUri:
            "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount",
        },
        dataSourceSetInfo: {
          datasourceType: "Microsoft.Storage/storageAccounts/blobServices",
          objectType: "DatasourceSet",
          resourceID:
            "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount",
          resourceLocation: "centraluseuap",
          resourceName: "blobstorageaccount",
          resourceType: "microsoft.storage/storageAccounts",
          resourceUri:
            "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount",
        },
        friendlyName: "blobstorageaccount\\blobbackupinstance",
        objectType: "BackupInstance",
        policyInfo: {
          policyId:
            "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupPolicies/blobpolicy",
          policyParameters: {
            backupDatasourceParametersList: [
              {
                autoProtectionSettings: {
                  enabled: true,
                  objectType: "BlobBackupRuleBasedAutoProtectionSettings",
                  rules: [
                    {
                      objectType: "BlobBackupAutoProtectionRule",
                      mode: "Exclude",
                      type: "Prefix",
                      pattern: "temp-",
                    },
                    {
                      objectType: "BlobBackupAutoProtectionRule",
                      mode: "Exclude",
                      type: "Prefix",
                      pattern: "test-",
                    },
                  ],
                },
                objectType: "BlobBackupDatasourceParametersForAutoProtection",
              },
            ],
          },
        },
      },
    },
  );
  console.log(result);
}
