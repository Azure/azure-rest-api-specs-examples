const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Remove Database principals permissions.
 *
 * @summary Remove Database principals permissions.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabaseRemovePrincipals.json
 */
async function kustoDatabaseRemovePrincipals() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase8";
  const databasePrincipalsToRemove = {
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
  const result = await client.databases.removePrincipals(
    resourceGroupName,
    clusterName,
    databaseName,
    databasePrincipalsToRemove,
  );
  console.log(result);
}
