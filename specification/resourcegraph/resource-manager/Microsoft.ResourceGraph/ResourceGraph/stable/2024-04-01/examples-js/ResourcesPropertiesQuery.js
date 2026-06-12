const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries the resources managed by Azure Resource Manager for scopes specified in the request.
 *
 * @summary queries the resources managed by Azure Resource Manager for scopes specified in the request.
 * x-ms-original-file: 2024-04-01/ResourcesPropertiesQuery.json
 */
async function accessAPropertiesField() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resources({
    query:
      "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | summarize count() by tostring(properties.storageProfile.osDisk.osType)",
    subscriptions: ["cfbbd179-59d2-4052-aa06-9270a38aa9d6"],
  });
  console.log(result);
}
