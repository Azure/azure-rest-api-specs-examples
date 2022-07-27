import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for Databases Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/PatchDatabase.json
     */
    /**
     * Sample code: Updates a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .update(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdb",
                new DatabaseUpdate()
                    .withSku(new Sku().withName("S1").withTier("Standard"))
                    .withCollation("SQL_Latin1_General_CP1_CI_AS")
                    .withMaxSizeBytes(1073741824L),
                Context.NONE);
    }
}
