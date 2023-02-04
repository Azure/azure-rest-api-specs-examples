const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Sentinel onboarding state
 *
 * @summary Create Sentinel onboarding state
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/onboardingStates/CreateSentinelOnboardingState.json
 */
async function createSentinelOnboardingState() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const sentinelOnboardingStateName = "default";
  const sentinelOnboardingStateParameter = {
    customerManagedKey: false,
  };
  const options = {
    sentinelOnboardingStateParameter,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.sentinelOnboardingStates.create(
    resourceGroupName,
    workspaceName,
    sentinelOnboardingStateName,
    options
  );
  console.log(result);
}
