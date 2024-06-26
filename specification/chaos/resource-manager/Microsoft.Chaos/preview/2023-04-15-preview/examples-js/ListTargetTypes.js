const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of Target Type resources for given location.
 *
 * @summary Get a list of Target Type resources for given location.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListTargetTypes.json
 */
async function listAllTargetTypesForWestus2Location() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const locationName = "westus2";
  const continuationToken = undefined;
  const options = { continuationToken };
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.targetTypes.list(locationName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
