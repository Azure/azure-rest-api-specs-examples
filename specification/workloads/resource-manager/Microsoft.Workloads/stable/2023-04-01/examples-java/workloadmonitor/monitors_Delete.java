/** Samples for Monitors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_Delete.json
     */
    /**
     * Sample code: Deletes a SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void deletesASAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.monitors().delete("myResourceGroup", "mySapMonitor", com.azure.core.util.Context.NONE);
    }
}
