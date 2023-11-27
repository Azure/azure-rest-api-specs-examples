const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Target Type resources for given location.
 *
 * @summary Get a Target Type resources for given location.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetTargetType.json
 */
async function getATargetTypeForWestus2Location() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const locationName = "westus2";
  const targetTypeName = "Microsoft-Agent";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.targetTypes.get(locationName, targetTypeName);
  console.log(result);
}
