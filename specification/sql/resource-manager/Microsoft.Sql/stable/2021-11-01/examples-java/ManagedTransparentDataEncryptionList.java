
import com.azure.core.util.Context;

/** Samples for ManagedDatabaseTransparentDataEncryption ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedTransparentDataEncryptionList.
     * json
     */
    /**
     * Sample code: Get a list of the database's transparent data encryptions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAListOfTheDatabaseSTransparentDataEncryptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseTransparentDataEncryptions()
            .listByDatabase("security-tde-resourcegroup", "securitytde", "testdb", Context.NONE);
    }
}
