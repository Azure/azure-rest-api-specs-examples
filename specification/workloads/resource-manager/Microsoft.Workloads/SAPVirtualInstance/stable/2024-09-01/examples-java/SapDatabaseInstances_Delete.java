
/**
 * Samples for SapDatabaseInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapDatabaseInstances_Delete.json
     */
    /**
     * Sample code: SAPDatabaseInstances_Delete.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPDatabaseInstancesDelete(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapDatabaseInstances().delete("test-rg", "X00", "databaseServer", com.azure.core.util.Context.NONE);
    }
}
