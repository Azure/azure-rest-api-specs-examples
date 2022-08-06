const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of reports within a given configuration profile assignment
 *
 * @summary Retrieve a list of reports within a given configuration profile assignment
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileAssignment.json
 */
async function listReportsByConfigurationProfilesAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const configurationProfileAssignmentName = "default";
  const vmName = "myVMName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reports.listByConfigurationProfileAssignments(
    resourceGroupName,
    configurationProfileAssignmentName,
    vmName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listReportsByConfigurationProfilesAssignment().catch(console.error);
