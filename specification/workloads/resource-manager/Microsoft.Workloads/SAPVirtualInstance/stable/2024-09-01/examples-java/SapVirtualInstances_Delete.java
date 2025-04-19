
/**
 * Samples for SapVirtualInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_Delete.json
     */
    /**
     * Sample code: SAPVirtualInstances_Delete.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesDelete(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().delete("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}
