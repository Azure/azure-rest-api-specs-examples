
/**
 * Samples for ManagedInstances ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceListByInstancePoolWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: List managed instances by instance pool with $expand=administrators/activedirectory.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedInstancesByInstancePoolWithExpandAdministratorsActivedirectory(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByInstancePool("Test1", "pool1", null,
            com.azure.core.util.Context.NONE);
    }
}
