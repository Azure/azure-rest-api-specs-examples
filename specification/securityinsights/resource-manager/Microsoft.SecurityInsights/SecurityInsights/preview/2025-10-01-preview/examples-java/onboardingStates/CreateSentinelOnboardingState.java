
/**
 * Samples for SentinelOnboardingStates Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/onboardingStates/CreateSentinelOnboardingState.json
     */
    /**
     * Sample code: Create Sentinel onboarding state.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createSentinelOnboardingState(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sentinelOnboardingStates().define("default").withExistingWorkspace("myRg", "myWorkspace")
            .withCustomerManagedKey(false).create();
    }
}
