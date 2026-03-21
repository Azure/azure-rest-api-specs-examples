const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a key vault in the specified subscription.
 *
 * @summary create or update a key vault in the specified subscription.
 * x-ms-original-file: 2026-02-01/createVaultWithNetworkAcls.json
 */
async function createOrUpdateAVaultWithNetworkAcls() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.createOrUpdate("sample-resource-group", "sample-vault", {
    location: "westus",
    properties: {
      enabledForDeployment: true,
      enabledForDiskEncryption: true,
      enabledForTemplateDeployment: true,
      networkAcls: {
        bypass: "AzureServices",
        defaultAction: "Deny",
        ipRules: [{ value: "124.56.78.91" }, { value: "'10.91.4.0/24'" }],
        virtualNetworkRules: [
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1",
          },
        ],
      },
      sku: { name: "standard", family: "A" },
      tenantId: "00000000-0000-0000-0000-000000000000",
    },
  });
  console.log(result);
}
