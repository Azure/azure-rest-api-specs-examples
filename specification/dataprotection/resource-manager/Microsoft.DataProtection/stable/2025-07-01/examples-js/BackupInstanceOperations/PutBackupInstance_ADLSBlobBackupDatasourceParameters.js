const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a backup instance in a backup vault
 *
 * @summary Create or update a backup instance in a backup vault
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/BackupInstanceOperations/PutBackupInstance_ADLSBlobBackupDatasourceParameters.json
 */
async function createBackupInstanceWithAdlsBlobBackupDatasourceParameters() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "54707983-993e-43de-8d94-074451394eda";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "adlsrg";
  const vaultName = "adlsvault";
  const backupInstanceName =
    "adlsstorageaccount-adlsstorageaccount-19a76f8a-c176-4f7d-819e-95157e2b0071";
  const parameters = {
    properties: {
      dataSourceInfo: {
        datasourceType: "Microsoft.Storage/storageAccounts/adlsBlobServices",
        objectType: "Datasource",
        resourceID:
          "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount",
        resourceLocation: "centraluseuap",
        resourceName: "adlsstorageaccount",
        resourceType: "microsoft.storage/storageAccounts",
        resourceUri:
          "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount",
      },
      dataSourceSetInfo: {
        datasourceType: "Microsoft.Storage/storageAccounts/adlsBlobServices",
        objectType: "DatasourceSet",
        resourceID:
          "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount",
        resourceLocation: "centraluseuap",
        resourceName: "adlsstorageaccount",
        resourceType: "microsoft.storage/storageAccounts",
        resourceUri:
          "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount",
      },
      friendlyName: "adlsstorageaccount\\adlsbackupinstance",
      objectType: "BackupInstance",
      policyInfo: {
        policyId:
          "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.DataProtection/backupVaults/adlsvault/backupPolicies/adlspolicy",
        policyParameters: {
          backupDatasourceParametersList: [
            {
              containersList: ["container1"],
              objectType: "AdlsBlobBackupDatasourceParameters",
            },
          ],
        },
      },
    },
    tags: { key1: "val1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vaultName,
    backupInstanceName,
    parameters,
  );
  console.log(result);
}
