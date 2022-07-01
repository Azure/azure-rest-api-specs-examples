import com.azure.core.util.Context;

/** Samples for ProviderInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/ProviderInstances_Get.json
     */
    /**
     * Sample code: Get properties of a SAP monitor Hana provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void getPropertiesOfASAPMonitorHanaProvider(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .getWithResponse("myResourceGroup", "mySapMonitor", "myProviderInstance", Context.NONE);
    }
}
