const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the lists of unsupported vm sizes for Service Fabric Managed Clusters.
 *
 * @summary Get the lists of unsupported vm sizes for Service Fabric Managed Clusters.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/managedUnsupportedVMSizesList_example.json
 */
async function listUnsupportedVMSizes() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedUnsupportedVMSizes.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
