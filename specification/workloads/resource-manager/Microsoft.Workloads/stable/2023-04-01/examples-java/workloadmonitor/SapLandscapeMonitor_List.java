/** Samples for SapLandscapeMonitor List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_List.json
     */
    /**
     * Sample code: Get properties of a SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void getPropertiesOfASAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapLandscapeMonitors()
            .listWithResponse("myResourceGroup", "mySapMonitor", com.azure.core.util.Context.NONE);
    }
}
