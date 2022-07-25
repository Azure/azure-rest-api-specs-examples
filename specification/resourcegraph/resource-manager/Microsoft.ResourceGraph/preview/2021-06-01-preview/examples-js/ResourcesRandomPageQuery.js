const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries the resources managed by Azure Resource Manager for scopes specified in the request.
 *
 * @summary Queries the resources managed by Azure Resource Manager for scopes specified in the request.
 * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesRandomPageQuery.json
 */
async function randomPageQuery() {
  const query = {
    options: { skip: 10, top: 2 },
    query: "Resources | where name contains 'test' | project id, name, type, location",
    subscriptions: ["cfbbd179-59d2-4052-aa06-9270a38aa9d6"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resources(query);
  console.log(result);
}

randomPageQuery().catch(console.error);
