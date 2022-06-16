const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function nfsV3AccountCreate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9101";
  const accountName = "sto4445";
  const parameters = {
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
    parameters
  );
  console.log(result);
}

nfsV3AccountCreate().catch(console.error);
