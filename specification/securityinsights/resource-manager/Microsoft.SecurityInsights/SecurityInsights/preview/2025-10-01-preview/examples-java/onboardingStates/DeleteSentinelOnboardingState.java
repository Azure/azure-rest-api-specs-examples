
/**
 * Samples for SentinelOnboardingStates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/onboardingStates/DeleteSentinelOnboardingState.json
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
