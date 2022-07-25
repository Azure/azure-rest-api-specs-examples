const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the status of a SQL pool operation
 *
 * @summary Get the status of a SQL pool operation
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetLocationHeaderResultWithSqlPool.json
 */
async function getTheResultOfAnOperationOnASqlAnalyticsPool() {
  const subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const sqlPoolName = "ExampleSqlPool";
  const operationId = "fedcba98-7654-4210-fedc-ba9876543210";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolOperationResults.getLocationHeaderResult(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    operationId
  );
  console.log(result);
}

getTheResultOfAnOperationOnASqlAnalyticsPool().catch(console.error);
