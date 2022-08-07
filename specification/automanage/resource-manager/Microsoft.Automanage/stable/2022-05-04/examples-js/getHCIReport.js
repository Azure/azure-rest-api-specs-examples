const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a report associated with a configuration profile assignment run
 *
 * @summary Get information about a report associated with a configuration profile assignment run
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCIReport.json
 */
async function getAReportForAHciConfigurationProfileAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const clusterName = "myClusterName";
  const configurationProfileAssignmentName = "default";
  const reportName = "b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.hCIReports.get(
    resourceGroupName,
    clusterName,
    configurationProfileAssignmentName,
    reportName
  );
  console.log(result);
}

getAReportForAHciConfigurationProfileAssignment().catch(console.error);
