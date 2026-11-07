
/**
 * Samples for VMInsights GetOnboardingStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-11-27-preview/getOnboardingStatusSingleVM.json
     */
    /**
     * Sample code: Get status for a VM that is actively reporting data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getStatusForAVMThatIsActivelyReportingData(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getVMInsights().getOnboardingStatusWithResponse(
            "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm",
            com.azure.core.util.Context.NONE);
    }
}
