
import com.azure.resourcemanager.sql.models.CreateDatabaseRestorePointDefinition;

/**
 * Samples for RestorePoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseRestorePointsPost.json
     */
    /**
     * Sample code: Creates datawarehouse database restore point.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsDatawarehouseDatabaseRestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorePoints().create("Default-SQL-SouthEastAsia",
            "testserver", "testDatabase", new CreateDatabaseRestorePointDefinition().withRestorePointLabel("mylabel"),
            com.azure.core.util.Context.NONE);
    }
}
