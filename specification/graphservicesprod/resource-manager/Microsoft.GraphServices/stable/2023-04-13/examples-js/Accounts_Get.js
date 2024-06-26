const { GraphServices } = require("@azure/arm-graphservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns account resource for a given name.
 *
 * @summary Returns account resource for a given name.
 * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Get.json
 */
async function getAccounts() {
  const subscriptionId =
    process.env["GRAPHSERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["GRAPHSERVICES_RESOURCE_GROUP"] || "testResourceGroupGRAM";
  const resourceName = "11111111-aaaa-1111-bbbb-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new GraphServices(credential, subscriptionId);
  const result = await client.accounts.get(resourceGroupName, resourceName);
  console.log(result);
}
