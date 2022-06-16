const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Allows modifying tags of environments. All other properties will be ignored.
 *
 * @summary Allows modifying tags of environments. All other properties will be ignored.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_Update.json
 */
async function environmentsUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const userName = "@me";
  const name = "{environmentName}";
  const dtlEnvironment = {
    tags: { tagName1: "tagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.environments.update(
    resourceGroupName,
    labName,
    userName,
    name,
    dtlEnvironment
  );
  console.log(result);
}

environmentsUpdate().catch(console.error);
