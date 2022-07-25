const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Maps Account.
 *
 * @summary Get a Maps Account.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/GetAccount.json
 */
async function getAccount() {
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myMapsAccount";
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.get(resourceGroupName, accountName);
  console.log(result);
}

getAccount().catch(console.error);
