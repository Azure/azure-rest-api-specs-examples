const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of contact profiles by Resource Group
 *
 * @summary Returns list of contact profiles by Resource Group
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfilesByResourceGroupList.json
 */
async function listOfContactProfilesByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.contactProfiles.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfContactProfilesByResourceGroup().catch(console.error);
