const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an Azure DevOps Repo.
 *
 * @summary Updates an Azure DevOps Repo.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoCreateOrUpdate.json
 */
async function azureDevOpsRepoCreateOrUpdate() {
  const subscriptionId =
    process.env["SECURITYDEVOPS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SECURITYDEVOPS_RESOURCE_GROUP"] || "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsOrgName = "myOrg";
  const azureDevOpsProjectName = "myProject";
  const azureDevOpsRepoName = "myRepo";
  const azureDevOpsRepo = {
    properties: {
      actionableRemediation: {
        branchConfiguration: { names: ["main"] },
        categories: ["Secrets"],
        severityLevels: ["High"],
        state: "Enabled",
      },
      repoId: "00000000-0000-0000-0000-000000000000",
      repoUrl: "https://dev.azure.com/myOrg/myProject/_git/myRepo",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsRepoOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsOrgName,
    azureDevOpsProjectName,
    azureDevOpsRepoName,
    azureDevOpsRepo
  );
  console.log(result);
}
