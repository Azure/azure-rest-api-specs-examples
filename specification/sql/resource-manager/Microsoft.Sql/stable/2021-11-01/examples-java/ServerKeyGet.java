
/**
 * Samples for ServerKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerKeyGet.json
     */
    /**
     * Sample code: Get the server key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheServerKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerKeys().getWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-4645", "someVault_someKey_01234567890123456789012345678901", com.azure.core.util.Context.NONE);
    }
}
