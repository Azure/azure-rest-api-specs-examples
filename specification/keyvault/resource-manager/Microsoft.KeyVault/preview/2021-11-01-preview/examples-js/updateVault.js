const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a key vault in the specified subscription.
 *
 * @summary Update a key vault in the specified subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/updateVault.json
 */
async function updateAnExistingVault() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-resource-group";
  const vaultName = "sample-vault";
  const parameters = {
    properties: {
      accessPolicies: [
        {
          objectId: "00000000-0000-0000-0000-000000000000",
          permissions: {
            certificates: [
              "get",
              "list",
              "delete",
              "create",
              "import",
              "update",
              "managecontacts",
              "getissuers",
              "listissuers",
              "setissuers",
              "deleteissuers",
              "manageissuers",
              "recover",
              "purge",
            ],
            keys: [
              "encrypt",
              "decrypt",
              "wrapKey",
              "unwrapKey",
              "sign",
              "verify",
              "get",
              "list",
              "create",
              "update",
              "import",
              "delete",
              "backup",
              "restore",
              "recover",
              "purge",
            ],
            secrets: ["get", "list", "set", "delete", "backup", "restore", "recover", "purge"],
          },
          tenantId: "00000000-0000-0000-0000-000000000000",
        },
      ],
      enabledForDeployment: true,
      enabledForDiskEncryption: true,
      enabledForTemplateDeployment: true,
      publicNetworkAccess: "Enabled",
      sku: { name: "standard", family: "A" },
      tenantId: "00000000-0000-0000-0000-000000000000",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.update(resourceGroupName, vaultName, parameters);
  console.log(result);
}

updateAnExistingVault().catch(console.error);
