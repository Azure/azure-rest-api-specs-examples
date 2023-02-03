const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a monitored GitHub repository.
 *
 * @summary Patch a monitored GitHub repository.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoUpdate.json
 */
async function gitHubRepoUpdate() {
  const subscriptionId =
    process.env["SECURITYDEVOPS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SECURITYDEVOPS_RESOURCE_GROUP"] || "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubOwnerName = "Azure";
  const gitHubRepoName = "azure-rest-api-specs";
  const gitHubRepo = {};
  const options = { gitHubRepo };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubRepoOperations.beginUpdateAndWait(
    resourceGroupName,
    gitHubConnectorName,
    gitHubOwnerName,
    gitHubRepoName,
    options
  );
  console.log(result);
}
