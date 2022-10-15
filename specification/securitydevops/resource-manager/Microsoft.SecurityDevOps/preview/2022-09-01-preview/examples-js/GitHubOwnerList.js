const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of monitored GitHub owners.
 *
 * @summary Returns a list of monitored GitHub owners.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerList.json
 */
async function gitHubOwnerList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const gitHubConnectorName = "testconnector";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gitHubOwnerOperations.list(
    resourceGroupName,
    gitHubConnectorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

gitHubOwnerList().catch(console.error);
