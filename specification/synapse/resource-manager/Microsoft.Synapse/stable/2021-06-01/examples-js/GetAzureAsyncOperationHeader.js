const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the status of an operation
 *
 * @summary Get the status of an operation
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAzureAsyncOperationHeader.json
 */
async function getAzureAsyncOperationHeaderResult() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "resourceGroup1";
  const workspaceName = "workspace1";
  const operationId = "01234567-89ab-4def-0123-456789abcdef";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.operations.getAzureAsyncHeaderResult(
    resourceGroupName,
    workspaceName,
    operationId
  );
  console.log(result);
}
