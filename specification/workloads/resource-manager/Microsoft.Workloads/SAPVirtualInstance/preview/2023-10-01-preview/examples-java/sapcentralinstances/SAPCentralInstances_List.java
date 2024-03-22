
/**
 * Samples for SapCentralInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapcentralinstances/SAPCentralInstances_List.json
     */
    /**
     * Sample code: SAPCentralInstances_List.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPCentralInstancesList(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralInstances().list("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
