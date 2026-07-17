const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves the VM Insights onboarding status for the specified resource or resource scope.
 *
 * @summary retrieves the VM Insights onboarding status for the specified resource or resource scope.
 * x-ms-original-file: 2018-11-27-preview/getOnboardingStatusSubscription.json
 */
async function getStatusForASubscriptionThatHasAtLeastOneVMThatIsActivelyReportingData() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.vMInsights.getOnboardingStatus(
    "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87",
  );
  console.log(result);
}
