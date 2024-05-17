const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a monitored GitHub owner.
 *
 * @summary Returns a monitored GitHub owner.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetGitHubOwners_example.json
 */
async function getGitHubOwners() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const ownerName = "myGitHubOwner";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.gitHubOwners.get(resourceGroupName, securityConnectorName, ownerName);
  console.log(result);
}
