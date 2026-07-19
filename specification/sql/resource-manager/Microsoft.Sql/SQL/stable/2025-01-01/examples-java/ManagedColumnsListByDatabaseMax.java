
import java.util.Arrays;

/**
 * Samples for ManagedDatabaseColumns ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedColumnsListByDatabaseMax.json
     */
    /**
     * Sample code: Filter managed database columns.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void filterManagedDatabaseColumns(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseColumns().listByDatabase("myRG", "serverName", "myDatabase",
            Arrays.asList("dbo"), Arrays.asList("customer", "address"), Arrays.asList("username"),
            Arrays.asList("schema asc", "table", "column desc"), null, com.azure.core.util.Context.NONE);
    }
}
