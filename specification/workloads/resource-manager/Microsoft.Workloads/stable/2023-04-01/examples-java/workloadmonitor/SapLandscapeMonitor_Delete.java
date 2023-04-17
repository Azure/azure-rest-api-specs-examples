/** Samples for SapLandscapeMonitor Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Delete.json
     */
    /**
     * Sample code: Deletes SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void deletesSAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapLandscapeMonitors()
            .deleteByResourceGroupWithResponse("myResourceGroup", "mySapMonitor", com.azure.core.util.Context.NONE);
    }
}
