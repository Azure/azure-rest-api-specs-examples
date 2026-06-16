const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric service resource with the specified name.
 *
 * @summary create or update a Service Fabric service resource with the specified name.
 * x-ms-original-file: 2026-03-01-preview/ServicePutOperation_example_min.json
 */
async function putAServiceWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.services.createOrUpdate("resRg", "myCluster", "myApp", "myService", {
    properties: {
      instanceCount: 1,
      partitionDescription: { partitionScheme: "Singleton" },
      serviceKind: "Stateless",
      serviceTypeName: "myServiceType",
    },
    tags: {},
  });
  console.log(result);
}
