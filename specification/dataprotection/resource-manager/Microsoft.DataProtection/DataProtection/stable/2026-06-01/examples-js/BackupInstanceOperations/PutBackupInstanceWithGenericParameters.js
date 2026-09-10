const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a backup instance in a backup vault
 *
 * @summary create or update a backup instance in a backup vault
 * x-ms-original-file: 2026-06-01/BackupInstanceOperations/PutBackupInstanceWithGenericParameters.json
 */
async function createBackupInstanceWithGenericBackupDatasourceParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "97cda027-4279-4cde-b4ff-19afa0021d87";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.createOrUpdate(
    "ESAN-ECYBVTRG",
    "ESANVault",
    "esan-volgroup-bi",
    {
      tags: { key1: "val1" },
      properties: {
        friendlyName: "esan-volgroup-bi",
        dataSourceInfo: {
          resourceID:
            "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/esan-volgroup",
          resourceUri: "SampleresourceUri123",
          datasourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceName: "esan-volgroup-bi",
          resourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceLocation: "eastus2euap",
          objectType: "Datasource",
        },
        dataSourceSetInfo: {
          resourceID:
            "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc",
          datasourceType: "Microsoft.ElasticSan/elasticSans/volumeGroups",
          resourceType: "Microsoft.ElasticSan/elasticSans",
          resourceLocation: "eastus2euap",
          objectType: "DatasourceSet",
        },
        policyInfo: {
          policyId:
            "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.DataProtection/backupVaults/ESANVault/backupPolicies/BVTPolicy",
          policyParameters: {
            dataStoreParametersList: [
              {
                resourceGroupId:
                  "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG",
                objectType: "AzureOperationalStoreParameters",
                dataStoreType: "OperationalStore",
              },
            ],
            backupDatasourceParametersList: [
              {
                objectType: "GenericBackupDatasourceParameters",
                resourceSelectors: ["vol1", "vol2", "vol3"],
              },
            ],
          },
        },
        objectType: "BackupInstance",
      },
    },
  );
  console.log(result);
}
