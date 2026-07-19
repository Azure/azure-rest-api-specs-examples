
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for SQL.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAPrivateLinkResourceForSQL(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getPrivateLinkResources().getWithResponse("Default", "test-svr", "plr",
            com.azure.core.util.Context.NONE);
    }
}
