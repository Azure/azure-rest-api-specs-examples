const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists available operations for the Kusto sub-resources inside Microsoft.Synapse provider.
 *
 * @summary Lists available operations for the Kusto sub-resources inside Microsoft.Synapse provider.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoOperationsList.json
 */
async function kustoOperationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoOperationsList().catch(console.error);
