
/**
 * Samples for ServerUsages ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerUsageList.json
     */
    /**
     * Sample code: List servers usages.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listServersUsages(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerUsages().listByServer("sqlcrudtest-6730", "sqlcrudtest-9007",
            com.azure.core.util.Context.NONE);
    }
}
