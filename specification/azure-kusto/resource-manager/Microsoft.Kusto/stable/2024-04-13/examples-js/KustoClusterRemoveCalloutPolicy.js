const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Removes callout policy for engine services.
 *
 * @summary Removes callout policy for engine services.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClusterRemoveCalloutPolicy.json
 */
async function kustoClusterDropCalloutPolicy() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const calloutPolicy = { calloutId: "*_kusto" };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginRemoveCalloutPolicyAndWait(
    resourceGroupName,
    clusterName,
    calloutPolicy,
  );
  console.log(result);
}
