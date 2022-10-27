const ServiceFabricManagementClient = require("@azure-rest/arm-servicefabric").default,
  { getLongRunningPoller } = require("@azure-rest/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric cluster resource with the specified name.
 *
 * @summary Create or update a Service Fabric cluster resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_min.json
 */
async function putAClusterWithMinimumParameters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const parameters = {
    body: {
      location: "eastus",
      properties: {
        diagnosticsStorageAccountConfig: {
          blobEndpoint: "https://diag.blob.core.windows.net/",
          protectedAccountKeyName: "StorageAccountKey1",
          queueEndpoint: "https://diag.queue.core.windows.net/",
          storageAccountName: "diag",
          tableEndpoint: "https://diag.table.core.windows.net/",
        },
        fabricSettings: [
          {
            name: "UpgradeService",
            parameters: [{ name: "AppPollIntervalInSeconds", value: "60" }],
          },
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
      },
      tags: {},
    },
  };
  const credential = new DefaultAzureCredential();
  const client = ServiceFabricManagementClient(credential);
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}",
      subscriptionId,
      resourceGroupName,
      clusterName
    )
    .put(parameters);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

putAClusterWithMinimumParameters().catch(console.error);
