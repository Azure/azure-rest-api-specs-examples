/** Samples for ProviderInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Delete.json
     */
    /**
     * Sample code: Deletes a SAP monitor provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void deletesASAPMonitorProvider(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .delete("myResourceGroup", "mySapMonitor", "myProviderInstance", com.azure.core.util.Context.NONE);
    }
}
