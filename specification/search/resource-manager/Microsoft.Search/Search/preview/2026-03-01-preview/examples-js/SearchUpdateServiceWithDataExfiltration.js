const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing search service in the given resource group.
 *
 * @summary updates an existing search service in the given resource group.
 * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceWithDataExfiltration.json
 */
async function searchUpdateServiceWithDataExfiltration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.update("rg1", "mysearchservice", {
    tags: { "app-name": "My e-commerce app", "new-tag": "Adding a new tag" },
    replicaCount: 2,
    dataExfiltrationProtections: ["BlockAll"],
  });
  console.log(result);
}
