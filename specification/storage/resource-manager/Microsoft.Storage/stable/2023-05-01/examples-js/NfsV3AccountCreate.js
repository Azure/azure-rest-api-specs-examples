const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary Asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/NfsV3AccountCreate.json
 */
async function nfsV3AccountCreate() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res9101";
  const accountName = "sto4445";
  const parameters = {
    enableExtendedGroups: true,
    isHnsEnabled: true,
    enableNfsV3: true,
    kind: "BlockBlobStorage",
    location: "eastus",
    networkRuleSet: {
      bypass: "AzureServices",
      defaultAction: "Allow",
      ipRules: [],
      virtualNetworkRules: [
        {
          virtualNetworkResourceId:
            "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Network/virtualNetworks/net123/subnets/subnet12",
        },
      ],
    },
    sku: { name: "Premium_LRS" },
    enableHttpsTrafficOnly: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginCreateAndWait(
    resourceGroupName,
    accountName,
    parameters,
  );
  console.log(result);
}
