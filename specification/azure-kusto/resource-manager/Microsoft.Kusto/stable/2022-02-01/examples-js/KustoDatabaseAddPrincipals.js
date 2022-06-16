const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add Database principals permissions.
 *
 * @summary Add Database principals permissions.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabaseAddPrincipals.json
 */
async function kustoDatabaseAddPrincipals() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase8";
  const databasePrincipalsToAdd = {
    value: [
      {
        name: "Some User",
        type: "User",
        appId: "",
        email: "user@microsoft.com",
        fqn: "aaduser=some_guid",
        role: "Admin",
      },
      {
        name: "Kusto",
        type: "Group",
        appId: "",
        email: "kusto@microsoft.com",
        fqn: "aadgroup=some_guid",
        role: "Viewer",
      },
      {
        name: "SomeApp",
        type: "App",
        appId: "some_guid_app_id",
        email: "",
        fqn: "aadapp=some_guid_app_id",
        role: "Admin",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.addPrincipals(
    resourceGroupName,
    clusterName,
    databaseName,
    databasePrincipalsToAdd
  );
  console.log(result);
}

kustoDatabaseAddPrincipals().catch(console.error);
