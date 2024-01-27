
import com.azure.core.util.Context;
import java.util.Arrays;

/** Samples for DatabaseColumns ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ColumnsListByDatabaseMax.json
     */
    /**
     * Sample code: Filter database columns.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void filterDatabaseColumns(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseColumns().listByDatabase("myRG", "serverName",
            "myDatabase", Arrays.asList("dbo"), Arrays.asList("customer", "address"), Arrays.asList("username"),
            Arrays.asList("schema asc", "table", "column desc"), null, Context.NONE);
    }
}
