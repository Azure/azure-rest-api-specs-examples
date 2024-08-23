
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDatabaseLedger.json
     */
    /**
     * Sample code: Creates a database with ledger on.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseWithLedgerOn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", new DatabaseInner().withLocation("southeastasia").withIsLedgerOn(true),
            com.azure.core.util.Context.NONE);
    }
}
