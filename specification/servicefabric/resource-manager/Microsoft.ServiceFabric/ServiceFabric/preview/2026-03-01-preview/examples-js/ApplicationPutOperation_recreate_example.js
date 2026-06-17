const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric application resource with the specified name.
 *
 * @summary create or update a Service Fabric application resource with the specified name.
 * x-ms-original-file: 2026-03-01-preview/ApplicationPutOperation_recreate_example.json
 */
async function putAnApplicationWithRecreateOption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applications.createOrUpdate("resRg", "myCluster", "myApp", {
    parameters: { param1: "value1" },
    typeName: "myAppType",
    typeVersion: "1.0",
    upgradePolicy: { recreateApplication: true },
    tags: {},
  });
  console.log(result);
}
