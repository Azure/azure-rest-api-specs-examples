
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDatabaseCopyMode.json
     */
    /**
     * Sample code: Creates a database as a copy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseAsACopy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "dbcopy",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("S0").withTier("Standard"))
                .withCreateMode(CreateMode.COPY).withSourceDatabaseId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
            com.azure.core.util.Context.NONE);
    }
}
