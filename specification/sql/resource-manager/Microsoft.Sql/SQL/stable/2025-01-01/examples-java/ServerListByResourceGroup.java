
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerListByResourceGroup.json
     */
    /**
     * Sample code: List servers by resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listServersByResourceGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().listByResourceGroup("sqlcrudtest-7398", null,
            com.azure.core.util.Context.NONE);
    }
}
