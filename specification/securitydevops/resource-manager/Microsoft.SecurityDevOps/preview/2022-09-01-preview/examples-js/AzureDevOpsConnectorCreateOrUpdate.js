const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Azure DevOps Connector.
 *
 * @summary Creates or updates an Azure DevOps Connector.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorCreateOrUpdate.json
 */
async function azureDevOpsConnectorCreateOrUpdate() {
  const subscriptionId =
    process.env["SECURITYDEVOPS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SECURITYDEVOPS_RESOURCE_GROUP"] || "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsConnector = {
    location: "West US",
    properties: {
      authorization: { code: "00000000000000000000" },
      orgs: [
        {
          name: "testOrg",
          projects: [{ name: "testProject", repos: ["testRepo"] }],
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsConnectorOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsConnector
  );
  console.log(result);
}
