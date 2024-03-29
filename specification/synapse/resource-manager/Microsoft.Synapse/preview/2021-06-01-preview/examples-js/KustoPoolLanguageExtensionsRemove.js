const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove a list of language extensions that can run within KQL queries.
 *
 * @summary Remove a list of language extensions that can run within KQL queries.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsRemove.json
 */
async function kustoPoolRemoveLanguageExtensions() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const languageExtensionsToRemove = {
    value: [{ languageExtensionName: "PYTHON" }, { languageExtensionName: "R" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.beginRemoveLanguageExtensionsAndWait(
    workspaceName,
    kustoPoolName,
    resourceGroupName,
    languageExtensionsToRemove
  );
  console.log(result);
}
