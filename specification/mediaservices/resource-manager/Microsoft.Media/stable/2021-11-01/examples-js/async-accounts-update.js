const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Media Services account
 *
 * @summary Updates an existing Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/async-accounts-update.json
 */
async function updateAMediaServicesAccounts() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const parameters = { tags: { key1: "value3" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.mediaservices.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

updateAMediaServicesAccounts().catch(console.error);
