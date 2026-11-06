
/**
 * Samples for VMInsights GetOnboardingStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-11-27-preview/getOnboardingStatusSubscription.json
     */
    /**
     * Sample code: Get status for a subscription that has at least one VM that is actively reporting data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getStatusForASubscriptionThatHasAtLeastOneVMThatIsActivelyReportingData(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getVMInsights().getOnboardingStatusWithResponse(
            "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87", com.azure.core.util.Context.NONE);
    }
}
