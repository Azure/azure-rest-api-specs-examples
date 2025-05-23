const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Azure Web Category.
 *
 * @summary Gets the specified Azure Web Category.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/AzureWebCategoryGet.json
 */
async function getAzureWebCategoryByName() {
  const subscriptionId =
    process.env["NETWORK_SUBSCRIPTION_ID"] || "4de8428a-4a92-4cea-90ff-b47128b8cab8";
  const name = "Arts";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.webCategories.get(name);
  console.log(result);
}
