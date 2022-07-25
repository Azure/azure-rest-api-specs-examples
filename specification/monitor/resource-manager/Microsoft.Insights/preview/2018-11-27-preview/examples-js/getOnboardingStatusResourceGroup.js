const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the VM Insights onboarding status for the specified resource or resource scope.
 *
 * @summary Retrieves the VM Insights onboarding status for the specified resource or resource scope.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusResourceGroup.json
 */
async function getStatusForAResourceGroupThatHasAtLeastOneVMThatIsActivelyReportingData() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/resource-group-with-vms";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.vMInsights.getOnboardingStatus(resourceUri);
  console.log(result);
}

getStatusForAResourceGroupThatHasAtLeastOneVMThatIsActivelyReportingData().catch(console.error);
