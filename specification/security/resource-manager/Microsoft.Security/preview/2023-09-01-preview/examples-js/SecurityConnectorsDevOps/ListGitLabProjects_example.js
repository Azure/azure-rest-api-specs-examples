const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of GitLab projects that are directly owned by given group and onboarded to the connector.
 *
 * @summary Gets a list of GitLab projects that are directly owned by given group and onboarded to the connector.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListGitLabProjects_example.json
 */
async function listGitLabProjects() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const groupFQName = "myGitLabGroup$mySubGroup";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gitLabProjects.list(
    resourceGroupName,
    securityConnectorName,
    groupFQName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
