
/**
 * Samples for NetworkSecurityPerimeterConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/NetworkSecurityPerimeterConfigurationsGet.json
     */
    /**
     * Sample code: Get an NSP config by name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAnNSPConfigByName(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().getWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-7398", "00000001-2222-3333-4444-111144444444.assoc1", com.azure.core.util.Context.NONE);
    }
}
