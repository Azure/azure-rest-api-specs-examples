
/**
 * Samples for SapCentralServerInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_Delete.json
     */
    /**
     * Sample code: SapCentralServerInstances_Delete.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sapCentralServerInstancesDelete(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().delete("test-rg", "X00", "centralServer", com.azure.core.util.Context.NONE);
    }
}
