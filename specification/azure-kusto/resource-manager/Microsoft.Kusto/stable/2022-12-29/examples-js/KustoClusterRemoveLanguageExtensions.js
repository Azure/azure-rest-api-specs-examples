const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove a list of language extensions that can run within KQL queries.
 *
 * @summary Remove a list of language extensions that can run within KQL queries.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClusterRemoveLanguageExtensions.json
 */
async function kustoClusterRemoveLanguageExtensions() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const languageExtensionsToRemove = {
    value: [{ languageExtensionName: "PYTHON" }, { languageExtensionName: "R" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginRemoveLanguageExtensionsAndWait(
    resourceGroupName,
    clusterName,
    languageExtensionsToRemove
  );
  console.log(result);
}
