const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get list of configuration profile assignments
 *
 * @summary Get list of configuration profile assignments
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsByMachineName.json
 */
async function listConfigurationProfileAssignmentsByResourceGroupAndMachine() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const machineName = "myMachineName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationProfileAssignments.listByMachineName(
    resourceGroupName,
    machineName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConfigurationProfileAssignmentsByResourceGroupAndMachine().catch(console.error);
