
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.SecondaryType;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDatabaseNamedReplica.json
     */
    /**
     * Sample code: Creates a database as named replica secondary.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseAsNamedReplicaSecondary(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia")
                .withSku(new Sku().withName("HS_Gen4").withTier("Hyperscale").withCapacity(2))
                .withCreateMode(CreateMode.SECONDARY)
                .withSourceDatabaseId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Sql/servers/testsvr1/databases/primarydb")
                .withSecondaryType(SecondaryType.NAMED),
            Context.NONE);
    }
}
