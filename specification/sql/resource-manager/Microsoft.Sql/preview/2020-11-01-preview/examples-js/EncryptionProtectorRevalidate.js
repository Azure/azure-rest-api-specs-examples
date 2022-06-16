const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Revalidates an existing encryption protector.
 *
 * @summary Revalidates an existing encryption protector.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/EncryptionProtectorRevalidate.json
 */
async function revalidatesTheEncryptionProtector() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const encryptionProtectorName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.encryptionProtectors.beginRevalidateAndWait(
    resourceGroupName,
    serverName,
    encryptionProtectorName
  );
  console.log(result);
}

revalidatesTheEncryptionProtector().catch(console.error);
