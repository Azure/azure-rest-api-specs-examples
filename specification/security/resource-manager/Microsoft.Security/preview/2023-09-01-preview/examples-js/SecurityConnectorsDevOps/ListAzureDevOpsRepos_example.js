const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of Azure DevOps repositories onboarded to the connector.
 *
 * @summary Returns a list of Azure DevOps repositories onboarded to the connector.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListAzureDevOpsRepos_example.json
 */
async function listAzureDevOpsRepos() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const orgName = "myAzDevOpsOrg";
  const projectName = "myAzDevOpsProject";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureDevOpsRepos.list(
    resourceGroupName,
    securityConnectorName,
    orgName,
    projectName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
