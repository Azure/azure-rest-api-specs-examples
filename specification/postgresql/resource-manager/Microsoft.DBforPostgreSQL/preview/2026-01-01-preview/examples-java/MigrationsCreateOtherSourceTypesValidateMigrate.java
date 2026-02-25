
import com.azure.resourcemanager.postgresqlflexibleserver.models.AdminCredentials;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationOption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationSecretParameters;
import com.azure.resourcemanager.postgresqlflexibleserver.models.OverwriteDatabasesOnTargetServer;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SourceType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SslMode;
import java.util.Arrays;

/**
 * Samples for Migrations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsCreateOtherSourceTypesValidateMigrate.json
     */
    /**
     * Sample code: Create a migration with other source type for validating and migrating.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createAMigrationWithOtherSourceTypeForValidatingAndMigrating(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().define("examplemigration").withRegion("eastus")
            .withExistingFlexibleServer("exampleresourcegroup", "exampleserver")
            .withMigrationMode(MigrationMode.OFFLINE).withMigrationOption(MigrationOption.VALIDATE_AND_MIGRATE)
            .withSourceType(SourceType.ON_PREMISES).withSslMode(SslMode.PREFER)
            .withSourceDbServerResourceId("examplesource:5432@exampleuser")
            .withSecretParameters(new MigrationSecretParameters().withAdminCredentials(new AdminCredentials()
                .withSourceServerPassword("fakeTokenPlaceholder").withTargetServerPassword("fakeTokenPlaceholder")))
            .withDbsToMigrate(
                Arrays.asList("exampledatabase1", "exampledatabase2", "exampledatabase3", "exampledatabase4"))
            .withOverwriteDbsInTarget(OverwriteDatabasesOnTargetServer.TRUE).create();
    }
}
