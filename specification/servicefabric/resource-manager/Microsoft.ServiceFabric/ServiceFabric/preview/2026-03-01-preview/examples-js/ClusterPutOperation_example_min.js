const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric cluster resource with the specified name.
 *
 * @summary create or update a Service Fabric cluster resource with the specified name.
 * x-ms-original-file: 2026-03-01-preview/ClusterPutOperation_example_min.json
 */
async function putAClusterWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusters.createOrUpdate("resRg", "myCluster", {
    location: "eastus",
    diagnosticsStorageAccountConfig: {
      blobEndpoint: "https://diag.blob.core.windows.net/",
      protectedAccountKeyName: "StorageAccountKey1",
      queueEndpoint: "https://diag.queue.core.windows.net/",
      storageAccountName: "diag",
      tableEndpoint: "https://diag.table.core.windows.net/",
    },
    fabricSettings: [
      { name: "UpgradeService", parameters: [{ name: "AppPollIntervalInSeconds", value: "60" }] },
    ],
    managementEndpoint: "http://myCluster.eastus.cloudapp.azure.com:19080",
    nodeTypes: [
      {
        name: "nt1vm",
        applicationPorts: { endPort: 30000, startPort: 20000 },
        clientConnectionEndpointPort: 19000,
        durabilityLevel: "Bronze",
        ephemeralPorts: { endPort: 64000, startPort: 49000 },
        httpGatewayEndpointPort: 19007,
        isPrimary: true,
        vmInstanceCount: 5,
      },
    ],
    reliabilityLevel: "Silver",
    upgradeMode: "Automatic",
    tags: {},
  });
  console.log(result);
}
