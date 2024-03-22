
/**
 * Samples for SapVirtualInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPVirtualInstances_Start.json
     */
    /**
     * Sample code: Start the SAP system.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheSAPSystem(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().start("test-rg", "X00", null, com.azure.core.util.Context.NONE);
    }
}
