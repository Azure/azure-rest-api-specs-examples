/** Samples for SapVirtualInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Start.json
     */
    /**
     * Sample code: SAPVirtualInstances_Start.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesStart(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().start("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
