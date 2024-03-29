const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SparkConfiguration by name in a workspace.
 *
 * @summary Get SparkConfiguration by name in a workspace.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/SparkConfiguration_Get.json
 */
async function getSparkConfigurationByName() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "exampleResourceGroup";
  const sparkConfigurationName = "exampleSparkConfigurationName";
  const workspaceName = "exampleWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sparkConfiguration.get(
    resourceGroupName,
    sparkConfigurationName,
    workspaceName
  );
  console.log(result);
}
