import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;
import java.time.OffsetDateTime;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/CreateDatabaseRestoreMode.json
     */
    /**
     * Sample code: Creates a database from restore with database deletion time.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseFromRestoreWithDatabaseDeletionTime(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "dbrestore",
                new DatabaseInner()
                    .withLocation("southeastasia")
                    .withSku(new Sku().withName("S0").withTier("Standard"))
                    .withCreateMode(CreateMode.RESTORE)
                    .withSourceDatabaseId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb")
                    .withSourceDatabaseDeletionDate(OffsetDateTime.parse("2017-07-14T06:41:06.613Z")),
                Context.NONE);
    }
}
