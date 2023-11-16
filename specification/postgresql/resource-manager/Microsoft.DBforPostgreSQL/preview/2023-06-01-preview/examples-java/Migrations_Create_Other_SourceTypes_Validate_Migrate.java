import com.azure.resourcemanager.postgresqlflexibleserver.models.AdminCredentials;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationOption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationSecretParameters;
import com.azure.resourcemanager.postgresqlflexibleserver.models.OverwriteDbsInTargetEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SourceType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SslMode;
import java.util.Arrays;

/** Samples for Migrations Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_Create_Other_SourceTypes_Validate_Migrate.json
     */
    /**
     * Sample code: Create Migration with other source types for Validate and Migrate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createMigrationWithOtherSourceTypesForValidateAndMigrate(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .migrations()
            .define("testmigration")
            .withRegion("westus")
            .withExistingFlexibleServer("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget")
            .withMigrationMode(MigrationMode.OFFLINE)
            .withMigrationOption(MigrationOption.VALIDATE_AND_MIGRATE)
            .withSourceType(SourceType.ON_PREMISES)
            .withSslMode(SslMode.PREFER)
            .withSourceDbServerResourceId("testsource:5432@pguser")
            .withSecretParameters(
                new MigrationSecretParameters()
                    .withAdminCredentials(
                        new AdminCredentials()
                            .withSourceServerPassword("fakeTokenPlaceholder")
                            .withTargetServerPassword("fakeTokenPlaceholder")))
            .withDbsToMigrate(Arrays.asList("db1", "db2", "db3", "db4"))
            .withOverwriteDbsInTarget(OverwriteDbsInTargetEnum.TRUE)
            .create();
    }
}
