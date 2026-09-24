
/**
 * Samples for ManagedInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceDelete.json
     */
    /**
     * Sample code: Delete managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().delete("testrg", "testinstance",
            com.azure.core.util.Context.NONE);
    }
}
