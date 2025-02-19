const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a dryrun job to do necessary check before actual creation
 *
 * @summary create a dryrun job to do necessary check before actual creation
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PutDryrun.json
 */
async function putDryrun() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const dryrunName = "dryrunName";
  const parameters = {
    parameters: {
      actionName: "createOrUpdate",
      authInfo: {
        name: "name",
        authType: "secret",
        secretInfo: { secretType: "rawValue", value: "secret" },
      },
      targetService: {
        type: "AzureResource",
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linkers.beginCreateDryrunAndWait(resourceUri, dryrunName, parameters);
  console.log(result);
}
