const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationVmImagePlan_example.json
 */
async function putNodeTypeWithVmImagePlan() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      dataDiskSizeGB: 200,
      isPrimary: false,
      vmImageOffer: "windows_2022_test",
      vmImagePlan: {
        name: "win_2022_test_20_10_gen2",
        product: "windows_2022_test",
        publisher: "testpublisher",
      },
      vmImagePublisher: "testpublisher",
      vmImageSku: "win_2022_test_20_10_gen2",
      vmImageVersion: "latest",
      vmInstanceCount: 10,
      vmSize: "Standard_D3",
    },
  });
  console.log(result);
}
