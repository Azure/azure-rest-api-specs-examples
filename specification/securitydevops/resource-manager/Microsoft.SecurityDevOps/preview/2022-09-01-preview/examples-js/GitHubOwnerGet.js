const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a monitored GitHub repository.
 *
 * @summary Returns a monitored GitHub repository.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerGet.json
 */
async function gitHubOwnerGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubOwnerName = "Azure";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubOwnerOperations.get(
    resourceGroupName,
    gitHubConnectorName,
    gitHubOwnerName
  );
  console.log(result);
}

gitHubOwnerGet().catch(console.error);
