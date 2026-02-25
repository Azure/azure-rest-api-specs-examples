
import com.azure.resourcemanager.postgresqlflexibleserver.models.AdminCredentials;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationOption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationSecretParameters;
import com.azure.resourcemanager.postgresqlflexibleserver.models.OverwriteDatabasesOnTargetServer;
import java.util.Arrays;

/**
 * Samples for Migrations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsCreateValidateOnly.json
     */
    /**
     * Sample code: Create a migration for validating only.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createAMigrationForValidatingOnly(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().define("examplemigration").withRegion("eastus")
            .withExistingFlexibleServer("exampleresourcegroup", "exampleserver")
            .withMigrationMode(MigrationMode.OFFLINE).withMigrationOption(MigrationOption.VALIDATE)
            .withSourceDbServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/servers/examplesource")
            .withSecretParameters(new MigrationSecretParameters().withAdminCredentials(new AdminCredentials()
                .withSourceServerPassword("fakeTokenPlaceholder").withTargetServerPassword("fakeTokenPlaceholder")))
            .withDbsToMigrate(
                Arrays.asList("exampledatabase1", "exampledatabase2", "exampledatabase3", "exampledatabase4"))
            .withOverwriteDbsInTarget(OverwriteDatabasesOnTargetServer.TRUE).create();
    }
}
