
import com.azure.resourcemanager.postgresqlflexibleserver.models.AdminCredentials;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationSecretParameters;
import com.azure.resourcemanager.postgresqlflexibleserver.models.OverwriteDatabasesOnTargetServer;
import java.util.Arrays;

/**
 * Samples for Migrations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsCreateWithPrivateEndpointServers.json
     */
    /**
     * Sample code: Create a migration with private endpoint.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createAMigrationWithPrivateEndpoint(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().define("examplemigration").withRegion("eastus")
            .withExistingFlexibleServer("exampleresourcegroup", "exampleserver")
            .withMigrationInstanceResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/flexibleServers/examplesourcemigration")
            .withMigrationMode(MigrationMode.OFFLINE)
            .withSourceDbServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/servers/examplesource")
            .withSecretParameters(new MigrationSecretParameters().withAdminCredentials(new AdminCredentials()
                .withSourceServerPassword("fakeTokenPlaceholder").withTargetServerPassword("fakeTokenPlaceholder")))
            .withDbsToMigrate(
                Arrays.asList("exampledatabase1", "exampledatabase2", "exampledatabase3", "exampledatabase4"))
            .withOverwriteDbsInTarget(OverwriteDatabasesOnTargetServer.TRUE).create();
    }
}
