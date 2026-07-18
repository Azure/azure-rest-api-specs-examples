
/**
 * Samples for TransparentDataEncryptions ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/TransparentDataEncryptionList.json
     */
    /**
     * Sample code: Get a list of the database's transparent data encryption.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAListOfTheDatabaseSTransparentDataEncryption(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTransparentDataEncryptions().listByDatabase("security-tde-resourcegroup",
            "securitytde", "testdb", com.azure.core.util.Context.NONE);
    }
}
