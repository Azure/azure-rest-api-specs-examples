
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationListFilter;

/**
 * Samples for Migrations ListByTargetServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * Migrations_ListByTargetServer.json
     */
    /**
     * Sample code: Migrations_ListByTargetServer.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        migrationsListByTargetServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().listByTargetServer("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget",
            MigrationListFilter.ALL, com.azure.core.util.Context.NONE);
    }
}
