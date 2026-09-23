
/**
 * Samples for ServerKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerKeyGetWithVersionlessKey.json
     */
    /**
     * Sample code: Get the server key with versionless key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheServerKeyWithVersionlessKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerKeys().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey", com.azure.core.util.Context.NONE);
    }
}
