const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 *
 * @summary updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 * x-ms-original-file: 2025-10-01-preview/UpdateAccount.json
 */
async function updateAccountTags() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.update("myResourceGroup", "myMapsAccount", {
    tags: { specialTag: "true" },
  });
  console.log(result);
}
