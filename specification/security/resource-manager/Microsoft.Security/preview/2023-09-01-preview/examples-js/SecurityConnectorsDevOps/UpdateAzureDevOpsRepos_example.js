const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a monitored Azure DevOps repository resource.
 *
 * @summary Updates a monitored Azure DevOps repository resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/UpdateAzureDevOpsRepos_example.json
 */
async function updateAzureDevOpsRepos() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const orgName = "myAzDevOpsOrg";
  const projectName = "myAzDevOpsProject";
  const repoName = "myAzDevOpsRepo";
  const azureDevOpsRepository = {
    properties: {
      actionableRemediation: { state: "Enabled" },
      onboardingState: "NotApplicable",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.azureDevOpsRepos.beginUpdateAndWait(
    resourceGroupName,
    securityConnectorName,
    orgName,
    projectName,
    repoName,
    azureDevOpsRepository,
  );
  console.log(result);
}
