import com.azure.core.util.Context;

/** Samples for SentinelOnboardingStates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/onboardingStates/DeleteSentinelOnboardingState.json
     */
    /**
     * Sample code: Delete Sentinel onboarding state.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteSentinelOnboardingState(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sentinelOnboardingStates().deleteWithResponse("myRg", "myWorkspace", "default", Context.NONE);
    }
}
