
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDatabaseDefaultMode.json
     */
    /**
     * Sample code: Creates a database with default mode.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADatabaseWithDefaultMode(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("S0").withTier("Standard"))
                .withCreateMode(CreateMode.DEFAULT).withCollation("SQL_Latin1_General_CP1_CI_AS")
                .withMaxSizeBytes(1073741824L),
            Context.NONE);
    }
}
