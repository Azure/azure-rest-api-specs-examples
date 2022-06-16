const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available workspaces under the specified resource group.
 *
 * @summary Lists all the available workspaces under the specified resource group.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_ListByResourceGroup.json
 */
async function getWorkspacesByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getWorkspacesByResourceGroup().catch(console.error);
