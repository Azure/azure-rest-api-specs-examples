const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to check the availability of name for resource.
 *
 * @summary check the availability of name for resource.
 * x-ms-original-file: 2026-01-01-preview/CheckScopedNameAvailability_DevCenterCatalog.json
 */
async function devcenterCatalogNameAvailability() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58fffff";
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.checkScopedNameAvailability.execute({
    name: "name1",
    type: "Microsoft.DevCenter/devcenters/catalogs",
    scope:
      "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso",
  });
  console.log(result);
}
