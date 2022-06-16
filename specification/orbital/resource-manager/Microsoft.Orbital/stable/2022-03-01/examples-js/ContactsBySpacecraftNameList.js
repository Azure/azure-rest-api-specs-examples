const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of contacts by spacecraftName
 *
 * @summary Returns list of contacts by spacecraftName
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactsBySpacecraftNameList.json
 */
async function listOfSpacecraft() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.contacts.list(resourceGroupName, spacecraftName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfSpacecraft().catch(console.error);
