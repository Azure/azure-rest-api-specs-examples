const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing encryption protector.
 *
 * @summary updates an existing encryption protector.
 * x-ms-original-file: 2025-01-01/EncryptionProtectorCreateOrUpdateServiceManaged.json
 */
async function updateTheEncryptionProtectorToServiceManaged() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.encryptionProtectors.createOrUpdate(
    "sqlcrudtest-7398",
    "sqlcrudtest-4645",
    "current",
    { serverKeyName: "ServiceManaged", serverKeyType: "ServiceManaged" },
  );
  console.log(result);
}
