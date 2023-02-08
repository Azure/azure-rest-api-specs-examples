/** Samples for SapVirtualInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Start.json
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
