const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update monitored AzureDevOps Connector details.
 *
 * @summary Update monitored AzureDevOps Connector details.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorUpdate.json
 */
async function azureDevOpsConnectorUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsConnector = {
    location: "West US",
    tags: { client: "dev-client", env: "dev" },
  };
  const options = {
    azureDevOpsConnector,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsConnectorOperations.beginUpdateAndWait(
    resourceGroupName,
    azureDevOpsConnectorName,
    options
  );
  console.log(result);
}

azureDevOpsConnectorUpdate().catch(console.error);
