const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete specified Update Run
 *
 * @summary Delete specified Update Run
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/DeleteUpdateRuns.json
 */
async function deleteAnUpdate() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "b8d594e5-51f3-4c11-9c54-a7771b81c712";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "testrg";
  const clusterName = "testcluster";
  const updateName = "Microsoft4.2203.2.32";
  const updateRunName = "23b779ba-0d52-4a80-8571-45ca74664ec3";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.updateRuns.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    updateName,
    updateRunName,
  );
  console.log(result);
}
