
/**
 * Samples for VMInsights GetOnboardingStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-11-27-preview/getOnboardingStatusResourceGroup.json
     */
    /**
     * Sample code: Get status for a resource group that has at least one VM that is actively reporting data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getStatusForAResourceGroupThatHasAtLeastOneVMThatIsActivelyReportingData(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getVMInsights().getOnboardingStatusWithResponse(
            "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/resource-group-with-vms",
            com.azure.core.util.Context.NONE);
    }
}
