const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Add a list of language extensions that can run within KQL queries.
 *
 * @summary Add a list of language extensions that can run within KQL queries.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClusterAddLanguageExtensions.json
 */
async function kustoClusterAddLanguageExtensions() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const languageExtensionsToAdd = {
    value: [{ languageExtensionName: "PYTHON" }, { languageExtensionName: "R" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginAddLanguageExtensionsAndWait(
    resourceGroupName,
    clusterName,
    languageExtensionsToAdd,
  );
  console.log(result);
}
