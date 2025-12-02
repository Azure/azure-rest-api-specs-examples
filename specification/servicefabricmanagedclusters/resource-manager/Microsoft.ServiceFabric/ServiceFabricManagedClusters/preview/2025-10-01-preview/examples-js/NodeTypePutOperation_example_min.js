const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperation_example_min.json
 */
async function putANodeTypeWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      dataDiskSizeGB: 200,
      isPrimary: false,
      vmImageOffer: "WindowsServer",
      vmImagePublisher: "MicrosoftWindowsServer",
      vmImageSku: "2016-Datacenter-Server-Core",
      vmImageVersion: "latest",
      vmInstanceCount: 10,
      vmSize: "Standard_D3",
    },
  });
  console.log(result);
}
