const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified private link resource.
 *
 * @summary Gets information about the specified private link resource.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateLinkResourceGet.json
 */
async function getPrivateLinkResource() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const privateLinkResourceName = "sampleacct";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResourceOperations.get(
    resourceGroupName,
    accountName,
    privateLinkResourceName
  );
  console.log(result);
}

getPrivateLinkResource().catch(console.error);
