const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries the resources managed by Azure Resource Manager for scopes specified in the request.
 *
 * @summary queries the resources managed by Azure Resource Manager for scopes specified in the request.
 * x-ms-original-file: 2024-04-01/ResourcesComplexQuery.json
 */
async function complexQuery() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resources({
    query:
      "Resources | project id, name, type, location | where type =~ 'Microsoft.Compute/virtualMachines' | summarize count() by location | top 3 by count_",
    subscriptions: ["cfbbd179-59d2-4052-aa06-9270a38aa9d6"],
  });
  console.log(result);
}
