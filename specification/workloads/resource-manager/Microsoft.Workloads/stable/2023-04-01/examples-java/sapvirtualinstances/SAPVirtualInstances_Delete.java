/** Samples for SapVirtualInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Delete.json
     */
    /**
     * Sample code: SAPVirtualInstances_Delete.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesDelete(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().delete("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
