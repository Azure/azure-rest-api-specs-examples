const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the Azure Front Door Standard, Azure Front Door Premium, and CDN profiles within a resource group.
 *
 * @summary Lists all of the Azure Front Door Standard, Azure Front Door Premium, and CDN profiles within a resource group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_ListByResourceGroup.json
 */
async function profilesListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.profiles.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

profilesListByResourceGroup().catch(console.error);
