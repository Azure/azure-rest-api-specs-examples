const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a backup instance in a backup vault
 *
 * @summary Create or update a backup instance in a backup vault
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/BackupInstanceOperations/PutBackupInstance_KubernetesClusterBackupDatasourceParameters.json
 */
async function createBackupInstanceWithKubernetesClusterBackupDatasourceParameters() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "62b829ee-7936-40c9-a1c9-47a93f9f3965";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "aksrg";
  const vaultName = "aksvault";
  const backupInstanceName = "aksbi";
  const parameters = {
    properties: {
      dataSourceInfo: {
        datasourceType: "Microsoft.ContainerService/managedclusters",
        objectType: "Datasource",
        resourceID:
          "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster",
        resourceLocation: "eastus2euap",
        resourceName: "akscluster",
        resourceType: "Microsoft.ContainerService/managedclusters",
        resourceUri:
          "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster",
      },
      dataSourceSetInfo: {
        datasourceType: "Microsoft.ContainerService/managedclusters",
        objectType: "DatasourceSet",
        resourceID:
          "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster",
        resourceLocation: "eastus2euap",
        resourceName: "akscluster",
        resourceType: "Microsoft.ContainerService/managedclusters",
        resourceUri:
          "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster",
      },
      friendlyName: "aksbi",
      objectType: "BackupInstance",
      policyInfo: {
        policyId:
          "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourcegroups/aksrg/providers/Microsoft.DataProtection/BackupVaults/aksvault/backupPolicies/akspolicy",
        policyParameters: {
          backupDatasourceParametersList: [
            {
              excludedNamespaces: ["kube-system"],
              excludedResourceTypes: ["v1/Secret"],
              includeClusterScopeResources: true,
              includedNamespaces: ["test"],
              includedResourceTypes: [],
              includedVolumeTypes: ["AzureDisk", "AzureFileShareSMB"],
              labelSelectors: [],
              objectType: "KubernetesClusterBackupDatasourceParameters",
              snapshotVolumes: true,
            },
          ],
          dataStoreParametersList: [
            {
              dataStoreType: "OperationalStore",
              objectType: "AzureOperationalStoreParameters",
              resourceGroupId:
                "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg",
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
