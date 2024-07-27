const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary Create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/NodeTypePutOperation_example_min.json
 */
async function putANodeTypeWithMinimumParameters() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const nodeTypeName = "BE";
  const parameters = {
    dataDiskSizeGB: 200,
    isPrimary: false,
    vmImageOffer: "WindowsServer",
    vmImagePublisher: "MicrosoftWindowsServer",
    vmImageSku: "2016-Datacenter-Server-Core",
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
