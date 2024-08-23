
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDatabaseMin.json
     */
    /**
     * Sample code: Creates a database with minimum number of parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsADatabaseWithMinimumNumberOfParameters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", new DatabaseInner().withLocation("southeastasia"), com.azure.core.util.Context.NONE);
    }
}
