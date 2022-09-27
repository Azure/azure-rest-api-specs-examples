/** Samples for SentinelOnboardingStates Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/onboardingStates/CreateSentinelOnboardingState.json
     */
    /**
     * Sample code: Create Sentinel onboarding state.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createSentinelOnboardingState(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .sentinelOnboardingStates()
            .define("default")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withCustomerManagedKey(false)
            .create();
    }
}
