const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of spacecrafts
 *
 * @summary Return list of spacecrafts
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftsByResourceGroupList.json
 */
async function listOfSpacecraftByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.spacecrafts.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfSpacecraftByResourceGroup().catch(console.error);
