const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.
 *
 * @summary Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListKeys.json
 */
async function listKeys() {
  const subscriptionId =
    process.env["MAPS_SUBSCRIPTION_ID"] || "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = process.env["MAPS_RESOURCE_GROUP"] || "myResourceGroup";
  const accountName = "myMapsAccount";
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.listKeys(resourceGroupName, accountName);
  console.log(result);
}
