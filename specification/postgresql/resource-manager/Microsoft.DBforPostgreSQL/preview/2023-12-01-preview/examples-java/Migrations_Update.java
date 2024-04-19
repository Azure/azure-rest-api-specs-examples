
import com.azure.resourcemanager.postgresqlflexibleserver.models.LogicalReplicationOnSourceDbEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationResource;

/**
 * Samples for Migrations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * Migrations_Update.json
     */
    /**
     * Sample code: Migrations_Update.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsUpdate(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        MigrationResource resource = manager.migrations().getWithResponse("ffffffff-ffff-ffff-ffff-ffffffffffff",
            "testrg", "testtarget", "testmigration", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSetupLogicalReplicationOnSourceDbIfNeeded(LogicalReplicationOnSourceDbEnum.TRUE).apply();
    }
}
