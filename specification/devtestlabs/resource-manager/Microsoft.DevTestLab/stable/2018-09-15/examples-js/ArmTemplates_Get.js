const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get azure resource manager template.
 *
 * @summary Get azure resource manager template.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArmTemplates_Get.json
 */
async function armTemplatesGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const artifactSourceName = "{artifactSourceName}";
  const name = "{armTemplateName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.armTemplates.get(
    resourceGroupName,
    labName,
    artifactSourceName,
    name
  );
  console.log(result);
}

armTemplatesGet().catch(console.error);
