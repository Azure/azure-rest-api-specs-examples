const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a monitored GitHub owner.
 *
 * @summary Create or update a monitored GitHub owner.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerCreateOrUpdate.json
 */
async function gitHubOwnerCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const gitHubOwnerName = "Azure";
  const gitHubOwner = {};
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.gitHubOwnerOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gitHubConnectorName,
    gitHubOwnerName,
    gitHubOwner
  );
  console.log(result);
}

gitHubOwnerCreateOrUpdate().catch(console.error);
