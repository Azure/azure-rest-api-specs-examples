const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing lab. This operation can take a while to complete.
 *
 * @summary Create or replace an existing lab. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateOrUpdate.json
 */
async function labsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const lab = {
    labStorageType: "{Standard|Premium}",
    location: "{location}",
    tags: { tagName1: "tagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.beginCreateOrUpdateAndWait(resourceGroupName, name, lab);
  console.log(result);
}

labsCreateOrUpdate().catch(console.error);
