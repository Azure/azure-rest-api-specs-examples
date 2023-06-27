import com.azure.resourcemanager.postgresqlflexibleserver.models.AdminCredentials;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationSecretParameters;
import com.azure.resourcemanager.postgresqlflexibleserver.models.OverwriteDbsInTargetEnum;
import java.util.Arrays;

/** Samples for Migrations Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/Migrations_Create.json
     */
    /**
     * Sample code: Migrations_Create.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsCreate(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .migrations()
            .define("testmigration")
            .withRegion("westus")
            .withExistingFlexibleServer("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget")
            .withMigrationMode(MigrationMode.OFFLINE)
            .withSourceDbServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource")
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
