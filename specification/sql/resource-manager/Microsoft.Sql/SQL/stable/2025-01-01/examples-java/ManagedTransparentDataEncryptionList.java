
/**
 * Samples for ManagedDatabaseTransparentDataEncryption ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedTransparentDataEncryptionList.json
     */
    /**
     * Sample code: Get a list of the database's transparent data encryptions.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAListOfTheDatabaseSTransparentDataEncryptions(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseTransparentDataEncryptions()
            .listByDatabase("security-tde-resourcegroup", "securitytde", "testdb", com.azure.core.util.Context.NONE);
    }
}
