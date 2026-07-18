
import com.azure.resourcemanager.sql.models.CreateDatabaseRestorePointDefinition;

/**
 * Samples for RestorePoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRestorePointsPost.json
     */
    /**
     * Sample code: Creates datawarehouse database restore point.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsDatawarehouseDatabaseRestorePoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorePoints().create("Default-SQL-SouthEastAsia", "testserver", "testDatabase",
            new CreateDatabaseRestorePointDefinition().withRestorePointLabel("mylabel"),
            com.azure.core.util.Context.NONE);
    }
}
