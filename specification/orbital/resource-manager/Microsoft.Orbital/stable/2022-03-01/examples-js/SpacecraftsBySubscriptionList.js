const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of spacecrafts
 *
 * @summary Return list of spacecrafts
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftsBySubscriptionList.json
 */
async function listOfSpacecraftBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.spacecrafts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfSpacecraftBySubscription().catch(console.error);
