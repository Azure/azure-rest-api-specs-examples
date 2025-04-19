
/**
 * Samples for SapApplicationServerInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_Delete.json
     */
    /**
     * Sample code: SAPApplicationServerInstances_Delete.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPApplicationServerInstancesDelete(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().delete("test-rg", "X00", "app01", com.azure.core.util.Context.NONE);
    }
}
