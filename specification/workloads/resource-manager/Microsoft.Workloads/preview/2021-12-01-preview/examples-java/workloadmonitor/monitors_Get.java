import com.azure.core.util.Context;

/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_Get.json
     */
    /**
     * Sample code: Get properties of a SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void getPropertiesOfASAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.monitors().getByResourceGroupWithResponse("myResourceGroup", "mySapMonitor", Context.NONE);
    }
}
