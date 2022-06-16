const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a workspace.
 *
 * @summary Gets the private link resources that need to be created for a workspace.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/PrivateLinkResourcesListByWorkspace.json
 */
async function workspacePrivateLinkResourcesListGroupIds() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspacePrivateLinkResources.listByWorkspace(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

workspacePrivateLinkResourcesListGroupIds().catch(console.error);
