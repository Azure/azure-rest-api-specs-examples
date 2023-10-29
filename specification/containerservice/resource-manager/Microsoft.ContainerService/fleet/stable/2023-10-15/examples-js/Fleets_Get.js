const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Fleet.
 *
 * @summary Gets a Fleet.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/Fleets_Get.json
 */
async function getsAFleetResource() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const fleetName = "fleet1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.get(resourceGroupName, fleetName);
  console.log(result);
}
