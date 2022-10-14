const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectList.json
 */
async function azureDevOpsProjectList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsOrgName = "myOrg";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureDevOpsProjectOperations.list(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsOrgName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

azureDevOpsProjectList().catch(console.error);
