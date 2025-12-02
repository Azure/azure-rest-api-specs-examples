const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric managed application resource with the specified name.
 *
 * @summary create or update a Service Fabric managed application resource with the specified name.
 * x-ms-original-file: 2025-10-01-preview/ApplicationPutOperation_example_min.json
 */
async function putAnApplicationWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.applications.createOrUpdate("resRg", "myCluster", "myApp", {
    location: "eastus",
    properties: {
      version:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0",
    },
  });
  console.log(result);
}
