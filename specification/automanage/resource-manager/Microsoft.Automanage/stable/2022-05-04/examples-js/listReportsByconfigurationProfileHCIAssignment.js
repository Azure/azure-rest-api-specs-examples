const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of reports within a given configuration profile assignment
 *
 * @summary Retrieve a list of reports within a given configuration profile assignment
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileHCIAssignment.json
 */
async function listReportsByHciConfigurationProfilesAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const clusterName = "myClusterName";
  const configurationProfileAssignmentName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hCIReports.listByConfigurationProfileAssignments(
    resourceGroupName,
    clusterName,
    configurationProfileAssignmentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listReportsByHciConfigurationProfilesAssignment().catch(console.error);
