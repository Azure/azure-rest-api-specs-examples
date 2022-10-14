const { MicrosoftSecurityDevOps } = require("@azure/arm-securitydevops");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update monitored AzureDevOps Org details.
 *
 * @summary Update monitored AzureDevOps Org details.
 * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgUpdate.json
 */
async function azureDevOpsOrgUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "westusrg";
  const azureDevOpsConnectorName = "testconnector";
  const azureDevOpsOrgName = "myOrg";
  const azureDevOpsOrg = {
    properties: { autoDiscovery: "Disabled" },
  };
  const options = { azureDevOpsOrg };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSecurityDevOps(credential, subscriptionId);
  const result = await client.azureDevOpsOrgOperations.beginUpdateAndWait(
    resourceGroupName,
    azureDevOpsConnectorName,
    azureDevOpsOrgName,
    options
  );
  console.log(result);
}

azureDevOpsOrgUpdate().catch(console.error);
