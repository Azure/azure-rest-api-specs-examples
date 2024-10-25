const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update Connector resource.
 *
 * @summary Create or update Connector resource.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PutConnector.json
 */
async function putConnector() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICELINKER_RESOURCE_GROUP"] || "test-rg";
  const location = "westus";
  const connectorName = "connectorName";
  const parameters = {
    authInfo: { authType: "secret" },
    secretStore: {
      keyVaultId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/test-kv",
    },
    targetService: {
      type: "AzureResource",
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.connector.beginCreateOrUpdateAndWait(
    subscriptionId,
    resourceGroupName,
    location,
    connectorName,
    parameters,
  );
  console.log(result);
}
