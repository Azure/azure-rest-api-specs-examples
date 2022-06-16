const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets status of an ongoing purge operation.
 *
 * @summary Gets status of an ongoing purge operation.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesPurgeOperation.json
 */
async function workspacePurgeOperation() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "aztest5048";
  const purgeId = "purge-970318e7-b859-4edb-8903-83b1b54d0b74";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.workspacePurge.getPurgeStatus(
    resourceGroupName,
    workspaceName,
    purgeId
  );
  console.log(result);
}

workspacePurgeOperation().catch(console.error);
