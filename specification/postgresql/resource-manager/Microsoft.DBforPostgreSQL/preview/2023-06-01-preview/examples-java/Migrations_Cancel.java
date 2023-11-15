import com.azure.resourcemanager.postgresqlflexibleserver.models.CancelEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MigrationResource;

/** Samples for Migrations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_Cancel.json
     */
    /**
     * Sample code: Cancel migration.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void cancelMigration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        MigrationResource resource =
            manager
                .migrations()
                .getWithResponse(
                    "ffffffff-ffff-ffff-ffff-ffffffffffff",
                    "testrg",
                    "testtarget",
                    "testmigration",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withCancel(CancelEnum.TRUE).apply();
    }
}
