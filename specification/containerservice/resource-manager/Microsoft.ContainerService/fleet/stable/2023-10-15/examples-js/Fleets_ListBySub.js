const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists fleets in the specified subscription.
 *
 * @summary Lists fleets in the specified subscription.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/Fleets_ListBySub.json
 */
async function listsTheFleetResourcesInASubscription() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fleets.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
