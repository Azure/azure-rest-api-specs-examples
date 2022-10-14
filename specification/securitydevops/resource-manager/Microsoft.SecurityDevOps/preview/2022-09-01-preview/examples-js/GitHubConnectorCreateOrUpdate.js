const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a monitored GitHub Connector resource.
 *
 * @summary Create or update a monitored GitHub Connector resource.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorCreateOrUpdate.json
 */
async function gitHubConnectorCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubConnector = {
    location: "West US",
    properties: { code: "00000000000000000000" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubConnectorOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gitHubConnectorName,
    gitHubConnector
  );
  console.log(result);
}

gitHubConnectorCreateOrUpdate().catch(console.error);
