
/**
 * Samples for ServerKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the server key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteTheServerKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerKeys().delete("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey_01234567890123456789012345678901", com.azure.core.util.Context.NONE);
    }
}
