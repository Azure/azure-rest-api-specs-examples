const { AstroManagementClient } = require("@azure/arm-astro");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the operations for the provider
 *
 * @summary List the operations for the provider
 * x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Operations_List_MaximumSet_Gen.json
 */
async function operationsListGeneratedByMaximumSetRule() {
  const subscriptionId =
    process.env["ASTRO_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AstroManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
