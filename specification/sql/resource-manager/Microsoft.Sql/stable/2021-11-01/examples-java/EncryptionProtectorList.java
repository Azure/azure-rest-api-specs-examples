
/**
 * Samples for EncryptionProtectors ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/EncryptionProtectorList.json
     */
    /**
     * Sample code: List encryption protectors by server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listEncryptionProtectorsByServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getEncryptionProtectors().listByServer("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
