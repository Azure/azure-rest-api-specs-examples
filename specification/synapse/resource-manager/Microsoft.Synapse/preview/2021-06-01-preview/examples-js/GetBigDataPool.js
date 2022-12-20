const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Big Data pool.
 *
 * @summary Get a Big Data pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/GetBigDataPool.json
 */
async function getABigDataPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const bigDataPoolName = "ExamplePool";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.bigDataPools.get(resourceGroupName, workspaceName, bigDataPoolName);
  console.log(result);
}
