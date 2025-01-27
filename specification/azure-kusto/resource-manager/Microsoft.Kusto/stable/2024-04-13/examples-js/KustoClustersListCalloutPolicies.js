const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns the allowed callout policies for the specified service.
 *
 * @summary Returns the allowed callout policies for the specified service.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClustersListCalloutPolicies.json
 */
async function kustoClusterListCalloutPolicies() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listCalloutPolicies(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
