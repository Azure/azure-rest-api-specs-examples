/** Samples for ProviderInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/NetWeaverProviderInstances_Get.json
     */
    /**
     * Sample code: Get properties of a SAP monitor NetWeaver provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void getPropertiesOfASAPMonitorNetWeaverProvider(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .getWithResponse("myResourceGroup", "mySapMonitor", "myProviderInstance", com.azure.core.util.Context.NONE);
    }
}
