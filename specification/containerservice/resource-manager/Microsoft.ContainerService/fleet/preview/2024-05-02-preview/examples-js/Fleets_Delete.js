const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Fleet
 *
 * @summary Delete a Fleet
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/Fleets_Delete.json
 */
async function deletesAFleetResourceAsynchronouslyWithALongRunningOperation() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const fleetName = "fleet1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.beginDeleteAndWait(resourceGroupName, fleetName);
  console.log(result);
}
