const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an outbound rule from the managed network of a machine learning workspace.
 *
 * @summary Gets an outbound rule from the managed network of a machine learning workspace.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/ManagedNetwork/getRule.json
 */
async function getManagedNetworkSettingsRule() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "aml-workspace-name";
  const ruleName = "name_of_the_fqdn_rule";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.managedNetworkSettingsRule.get(
    resourceGroupName,
    workspaceName,
    ruleName,
  );
  console.log(result);
}
