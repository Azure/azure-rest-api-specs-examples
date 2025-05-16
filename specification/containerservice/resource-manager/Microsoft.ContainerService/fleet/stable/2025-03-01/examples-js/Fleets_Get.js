const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a Fleet.
 *
 * @summary gets a Fleet.
 * x-ms-original-file: 2025-03-01/Fleets_Get.json
 */
async function getsAFleetResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.get("rg1", "fleet1");
  console.log(result);
}
