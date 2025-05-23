
/**
 * Samples for SentinelOnboardingStates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * onboardingStates/DeleteSentinelOnboardingState.json
     */
    /**
     * Sample code: Delete Sentinel onboarding state.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteSentinelOnboardingState(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sentinelOnboardingStates().deleteWithResponse("myRg", "myWorkspace", "default",
            com.azure.core.util.Context.NONE);
    }
}
