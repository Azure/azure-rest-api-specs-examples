const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric managed service resource with the specified name.
 *
 * @summary create or update a Service Fabric managed service resource with the specified name.
 * x-ms-original-file: 2025-10-01-preview/ServicePutOperation_example_min.json
 */
async function putAServiceWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.services.createOrUpdate("resRg", "myCluster", "myApp", "myService", {
    location: "eastus",
    properties: {
      instanceCount: 1,
      partitionDescription: { partitionScheme: "Singleton" },
      serviceKind: "Stateless",
      serviceTypeName: "myServiceType",
    },
  });
  console.log(result);
}
