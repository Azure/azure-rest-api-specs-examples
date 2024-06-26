const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate either the primary or secondary key for use with the Maps APIs. The old key will stop working immediately.
 *
 * @summary Regenerate either the primary or secondary key for use with the Maps APIs. The old key will stop working immediately.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/RegenerateKey.json
 */
async function regenerateKey() {
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myMapsAccount";
  const keySpecification = { keyType: "primary" };
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.regenerateKeys(
    resourceGroupName,
    accountName,
    keySpecification
  );
  console.log(result);
}

regenerateKey().catch(console.error);
