const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Allows modifying tags of user profiles. All other properties will be ignored.
 *
 * @summary Allows modifying tags of user profiles. All other properties will be ignored.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_Update.json
 */
async function usersUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{devtestlabName}";
  const name = "{userName}";
  const user = { tags: { tagName1: "tagValue1" } };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.users.update(resourceGroupName, labName, name, user);
  console.log(result);
}

usersUpdate().catch(console.error);
