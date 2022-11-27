const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List applicationGroups.
 *
 * @summary List applicationGroups.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_ListByResourceGroup.json
 */
async function applicationGroupListByResourceGroup() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = "resourceGroup1";
  const filter = "applicationGroupType eq 'RailApplication'";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationGroups.listByResourceGroup(resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

applicationGroupListByResourceGroup().catch(console.error);
