const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an existing Linker.
 *
 * @summary Operation to update an existing Linker.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PatchLinker.json
 */
async function patchLinker() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const linkerName = "linkName";
  const parameters = {
    authInfo: {
      authType: "servicePrincipalSecret",
      clientId: "name",
      principalId: "id",
      secret: "secret",
    },
    targetService: {
      type: "AzureResource",
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linker.beginUpdateAndWait(resourceUri, linkerName, parameters);
  console.log(result);
}
