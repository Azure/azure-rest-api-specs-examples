/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_List.json
     */
    /**
     * Sample code: List all SAP monitors in a subscription.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void listAllSAPMonitorsInASubscription(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
