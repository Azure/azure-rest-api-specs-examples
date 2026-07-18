
/**
 * Samples for ServerKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerKeyGet.json
     */
    /**
     * Sample code: Get the server key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheServerKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerKeys().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey_01234567890123456789012345678901", com.azure.core.util.Context.NONE);
    }
}
