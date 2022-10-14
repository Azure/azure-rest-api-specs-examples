const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a monitored GitHub repository.
 *
 * @summary Returns a monitored GitHub repository.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoGet.json
 */
async function gitHubRepoGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubOwnerName = "Azure";
  const gitHubRepoName = "39093389";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubRepoOperations.get(
    resourceGroupName,
    gitHubConnectorName,
    gitHubOwnerName,
    gitHubRepoName
  );
  console.log(result);
}

gitHubRepoGet().catch(console.error);
