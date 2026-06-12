const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries the resources managed by Azure Resource Manager for scopes specified in the request.
 *
 * @summary queries the resources managed by Azure Resource Manager for scopes specified in the request.
 * x-ms-original-file: 2024-04-01/ResourcesFacetQuery.json
 */
async function queryWithAFacetRequest() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resources({
    facets: [
      { expression: "location", options: { top: 3, sortOrder: "desc" } },
      {
        expression: "properties.storageProfile.osDisk.osType",
        options: { top: 3, sortOrder: "desc" },
      },
      { expression: "nonExistingColumn", options: { top: 3, sortOrder: "desc" } },
      {
        expression: "resourceGroup",
        options: { top: 3, sortBy: "tolower(resourceGroup)", sortOrder: "asc" },
      },
      { expression: "resourceGroup", options: { top: 3, filter: "resourceGroup contains 'test'" } },
    ],
    query:
      "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project id, name, location, resourceGroup, properties.storageProfile.osDisk.osType | limit 5",
    subscriptions: ["cfbbd179-59d2-4052-aa06-9270a38aa9d6"],
  });
  console.log(result);
}
