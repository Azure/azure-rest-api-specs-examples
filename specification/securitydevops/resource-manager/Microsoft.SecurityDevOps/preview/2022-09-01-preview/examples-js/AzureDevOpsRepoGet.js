const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a monitored AzureDevOps Project resource for a given ID.
 *
 * @summary Returns a monitored AzureDevOps Project resource for a given ID.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoGet.json
 */
async function azureDevOpsRepoGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsOrgName = "myOrg";
  const azureDevOpsProjectName = "myProject";
  const azureDevOpsRepoName = "myRepo";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsRepoOperations.get(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsOrgName,
    azureDevOpsProjectName,
    azureDevOpsRepoName
  );
  console.log(result);
}

azureDevOpsRepoGet().catch(console.error);
