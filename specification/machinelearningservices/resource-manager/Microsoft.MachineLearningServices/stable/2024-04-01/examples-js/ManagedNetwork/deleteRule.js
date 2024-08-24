const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an outbound rule from the managed network of a machine learning workspace.
 *
 * @summary Deletes an outbound rule from the managed network of a machine learning workspace.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/ManagedNetwork/deleteRule.json
 */
async function deleteManagedNetworkSettingsRule() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "aml-workspace-name";
  const ruleName = "rule-name";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.managedNetworkSettingsRule.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    ruleName,
  );
  console.log(result);
}
