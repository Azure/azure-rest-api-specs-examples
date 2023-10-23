const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the private link resources available for this workspace.
 *
 * @summary List the private link resources available for this workspace.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateLinkResources_ListByWorkspace.json
 */
async function privateLinkResourcesListByWorkspace() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByWorkspace(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
