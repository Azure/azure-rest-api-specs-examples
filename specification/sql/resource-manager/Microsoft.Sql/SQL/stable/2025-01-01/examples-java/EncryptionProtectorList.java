
/**
 * Samples for EncryptionProtectors ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EncryptionProtectorList.json
     */
    /**
     * Sample code: List encryption protectors by server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listEncryptionProtectorsByServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEncryptionProtectors().listByServer("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
