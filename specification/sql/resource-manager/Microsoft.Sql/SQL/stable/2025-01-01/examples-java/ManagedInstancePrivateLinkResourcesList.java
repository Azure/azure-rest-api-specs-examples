
/**
 * Samples for ManagedInstancePrivateLinkResources ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstancePrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for SQL.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsPrivateLinkResourcesForSQL(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstancePrivateLinkResources().listByManagedInstance("Default", "test-cl",
            com.azure.core.util.Context.NONE);
    }
}
