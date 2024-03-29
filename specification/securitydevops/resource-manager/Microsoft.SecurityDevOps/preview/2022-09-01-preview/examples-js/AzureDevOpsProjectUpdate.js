const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update monitored AzureDevOps Project details.
 *
 * @summary Update monitored AzureDevOps Project details.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectUpdate.json
 */
async function azureDevOpsProjectUpdate() {
  const subscriptionId =
    process.env["SECURITYDEVOPS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SECURITYDEVOPS_RESOURCE_GROUP"] || "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsOrgName = "myOrg";
  const azureDevOpsProjectName = "myProject";
  const azureDevOpsProject = {
    properties: { autoDiscovery: "Disabled" },
  };
  const options = {
    azureDevOpsProject,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsProjectOperations.beginUpdateAndWait(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsOrgName,
    azureDevOpsProjectName,
    options
  );
  console.log(result);
}
