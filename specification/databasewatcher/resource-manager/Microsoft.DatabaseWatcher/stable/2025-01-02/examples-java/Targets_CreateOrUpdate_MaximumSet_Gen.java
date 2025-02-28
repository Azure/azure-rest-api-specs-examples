
import com.azure.resourcemanager.databasewatcher.models.SqlDbSingleDatabaseTargetProperties;
import com.azure.resourcemanager.databasewatcher.models.TargetAuthenticationType;

/**
 * Samples for Targets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Targets_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        targetsCreateOrUpdateMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.targets().define("monitoringh22eed").withExistingWatcher("apiTest-ddat4p", "databasemo3ej9ih")
            .withProperties(new SqlDbSingleDatabaseTargetProperties()
                .withTargetAuthenticationType(TargetAuthenticationType.AAD).withConnectionServerName("sqlServero1ihe2")
                .withSqlDbResourceId(
                    "/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"))
            .create();
    }
}
