const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary Create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/NodeTypePutOperationVmImagePlan_example.json
 */
async function putNodeTypeWithVMImagePlan() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const nodeTypeName = "BE";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    nodeTypeName,
    parameters,
  );
  console.log(result);
}
