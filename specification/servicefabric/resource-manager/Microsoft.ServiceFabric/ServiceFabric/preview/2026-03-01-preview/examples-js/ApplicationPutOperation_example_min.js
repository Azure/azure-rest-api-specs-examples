const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric application resource with the specified name.
 *
 * @summary create or update a Service Fabric application resource with the specified name.
 * x-ms-original-file: 2026-03-01-preview/ApplicationPutOperation_example_min.json
 */
async function putAnApplicationWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applications.createOrUpdate("resRg", "myCluster", "myApp", {
    location: "eastus",
    removeApplicationCapacity: false,
    typeName: "myAppType",
    typeVersion: "1.0",
    tags: {},
  });
  console.log(result);
}
