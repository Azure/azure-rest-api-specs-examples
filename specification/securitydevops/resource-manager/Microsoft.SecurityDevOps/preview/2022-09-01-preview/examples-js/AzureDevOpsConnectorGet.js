const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a monitored AzureDevOps Connector resource for a given ID.
 *
 * @summary Returns a monitored AzureDevOps Connector resource for a given ID.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorGet.json
 */
async function azureDevOpsConnectorGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsConnectorOperations.get(
    resourceGroupName,
    azureDevOpsConnectorName
  );
  console.log(result);
}

azureDevOpsConnectorGet().catch(console.error);
