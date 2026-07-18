
/**
 * Samples for ManagedInstancePrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstancePrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for SQL.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAPrivateLinkResourceForSQL(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstancePrivateLinkResources().getWithResponse("Default", "test-cl", "plr",
            com.azure.core.util.Context.NONE);
    }
}
