
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationListFilter;

/**
 * Samples for Migrations ListByTargetServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsListByTargetServer.json
     */
    /**
     * Sample code: List all migrations of a target flexible server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllMigrationsOfATargetFlexibleServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().listByTargetServer("exampleresourcegroup", "exampleserver", MigrationListFilter.ALL,
            com.azure.core.util.Context.NONE);
    }
}
