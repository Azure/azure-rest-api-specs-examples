const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the linked services instances in a workspace.
 *
 * @summary Gets the linked services instances in a workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedServicesListByWorkspace.json
 */
async function linkedServicesListByWorkspace() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "mms-eus";
  const workspaceName = "TestLinkWS";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.linkedServices.listByWorkspace(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

linkedServicesListByWorkspace().catch(console.error);
