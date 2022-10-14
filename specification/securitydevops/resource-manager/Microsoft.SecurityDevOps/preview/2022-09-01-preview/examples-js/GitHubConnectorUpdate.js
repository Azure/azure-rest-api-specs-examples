const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update monitored GitHub Connector details.
 *
 * @summary Update monitored GitHub Connector details.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorUpdate.json
 */
async function gitHubConnectorUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubConnector = {
    location: "West US",
    tags: { client: "dev-client", env: "dev" },
  };
  const options = { gitHubConnector };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubConnectorOperations.beginUpdateAndWait(
    resourceGroupName,
    gitHubConnectorName,
    options
  );
  console.log(result);
}

gitHubConnectorUpdate().catch(console.error);
