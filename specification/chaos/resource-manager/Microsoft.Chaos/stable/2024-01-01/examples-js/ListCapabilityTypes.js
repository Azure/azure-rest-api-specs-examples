const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of Capability Type resources for given Target Type and location.
 *
 * @summary Get a list of Capability Type resources for given Target Type and location.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/ListCapabilityTypes.json
 */
async function listAllCapabilityTypesForAVirtualMachineTargetResourceOnWestus2Location() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const locationName = "westus2";
  const targetTypeName = "Microsoft-VirtualMachine";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.capabilityTypes.list(locationName, targetTypeName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
