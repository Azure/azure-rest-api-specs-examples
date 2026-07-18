
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerGetWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: Get server with $expand=administrators/activedirectory.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getServerWithExpandAdministratorsActivedirectory(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().getByResourceGroupWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            null, com.azure.core.util.Context.NONE);
    }
}
