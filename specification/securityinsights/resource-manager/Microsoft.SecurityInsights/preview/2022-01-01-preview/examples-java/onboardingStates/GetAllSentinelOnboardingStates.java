import com.azure.core.util.Context;

/** Samples for SentinelOnboardingStates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/onboardingStates/GetAllSentinelOnboardingStates.json
     */
    /**
     * Sample code: Get all Sentinel onboarding states.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSentinelOnboardingStates(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sentinelOnboardingStates().listWithResponse("myRg", "myWorkspace", Context.NONE);
    }
}
