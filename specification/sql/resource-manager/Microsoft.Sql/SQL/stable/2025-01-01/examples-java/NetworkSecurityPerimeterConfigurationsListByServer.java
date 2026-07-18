
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/NetworkSecurityPerimeterConfigurationsListByServer.json
     */
    /**
     * Sample code: List NSP configs by server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listNSPConfigsByServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().listByServer("sqlcrudtest-7398",
            "sqlcrudtest-7398", com.azure.core.util.Context.NONE);
    }
}
