/** Samples for ProviderInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/MsSqlServerProviderInstance_Get.json
     */
    /**
     * Sample code: Get properties of a MsSqlServer provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void getPropertiesOfAMsSqlServerProvider(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .getWithResponse("myResourceGroup", "mySapMonitor", "myProviderInstance", com.azure.core.util.Context.NONE);
    }
}
