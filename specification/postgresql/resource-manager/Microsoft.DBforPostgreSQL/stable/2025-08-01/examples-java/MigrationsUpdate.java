
import com.azure.resourcemanager.postgresqlflexibleserver.models.LogicalReplicationOnSourceServer;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Migration;

/**
 * Samples for Migrations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/MigrationsUpdate.
     * json
     */
    /**
     * Sample code: Update an existing migration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        updateAnExistingMigration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Migration resource = manager.migrations().getWithResponse("exampleresourcegroup", "exampleserver",
            "examplemigration", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSetupLogicalReplicationOnSourceDbIfNeeded(LogicalReplicationOnSourceServer.TRUE).apply();
    }
}
