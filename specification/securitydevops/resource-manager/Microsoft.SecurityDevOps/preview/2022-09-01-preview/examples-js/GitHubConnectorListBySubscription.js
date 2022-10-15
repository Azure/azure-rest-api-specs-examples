const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of monitored GitHub Connectors.
 *
 * @summary Returns a list of monitored GitHub Connectors.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorListBySubscription.json
 */
async function gitHubConnectorListBySubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gitHubConnectorOperations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

gitHubConnectorListBySubscription().catch(console.error);
