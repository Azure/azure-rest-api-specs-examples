
/**
 * Samples for SentinelOnboardingStates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/onboardingStates/GetSentinelOnboardingState.json
     */
    /**
     * Sample code: Get Sentinel onboarding state.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getSentinelOnboardingState(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sentinelOnboardingStates().getWithResponse("myRg", "myWorkspace", "default",
            com.azure.core.util.Context.NONE);
    }
}
