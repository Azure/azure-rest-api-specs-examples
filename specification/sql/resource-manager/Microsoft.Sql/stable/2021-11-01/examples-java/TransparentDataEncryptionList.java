
/**
 * Samples for TransparentDataEncryptions ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/TransparentDataEncryptionList.json
     */
    /**
     * Sample code: Get a list of the database's transparent data encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAListOfTheDatabaseSTransparentDataEncryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getTransparentDataEncryptions()
            .listByDatabase("security-tde-resourcegroup", "securitytde", "testdb", com.azure.core.util.Context.NONE);
    }
}
