const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the existing origins within an origin group.
 *
 * @summary Lists all of the existing origins within an origin group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_ListByOriginGroup.json
 */
async function afdOriginsListByOriginGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdOrigins.listByOriginGroup(
    resourceGroupName,
    profileName,
    originGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

afdOriginsListByOriginGroup().catch(console.error);
