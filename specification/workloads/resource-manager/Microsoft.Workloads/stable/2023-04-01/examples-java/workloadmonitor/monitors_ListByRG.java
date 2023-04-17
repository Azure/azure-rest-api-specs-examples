/** Samples for Monitors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_ListByRG.json
     */
    /**
     * Sample code: List all SAP monitors in a resource group.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void listAllSAPMonitorsInAResourceGroup(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.monitors().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
