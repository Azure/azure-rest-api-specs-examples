const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of monitored GitHub repositories.
 *
 * @summary Returns a list of monitored GitHub repositories.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoListByConnector.json
 */
async function gitHubRepoListByConnector() {
  const subscriptionId =
    process.env["SECURITYDEVOPS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SECURITYDEVOPS_RESOURCE_GROUP"] || "westusrg";
  const gitHubConnectorName = "testconnector";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gitHubRepoOperations.listByConnector(
    resourceGroupName,
    gitHubConnectorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
