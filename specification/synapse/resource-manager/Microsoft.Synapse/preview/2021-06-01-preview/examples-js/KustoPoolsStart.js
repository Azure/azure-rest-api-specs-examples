const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts a Kusto pool.
 *
 * @summary Starts a Kusto pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStart.json
 */
async function kustoPoolsStop() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.beginStartAndWait(
    workspaceName,
    kustoPoolName,
    resourceGroupName
  );
  console.log(result);
}
