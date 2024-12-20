const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update Linker resource.
 *
 * @summary Create or update Linker resource.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PutLinker.json
 */
async function putLinker() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const linkerName = "linkName";
  const parameters = {
    authInfo: {
      name: "name",
      authType: "secret",
      secretInfo: { secretType: "rawValue", value: "secret" },
    },
    targetService: {
      type: "AzureResource",
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db",
    },
    vNetSolution: { type: "serviceEndpoint" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linker.beginCreateOrUpdateAndWait(
    resourceUri,
    linkerName,
    parameters,
  );
  console.log(result);
}
