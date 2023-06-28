import com.azure.resourcemanager.postgresqlflexibleserver.fluent.models.MigrationNameAvailabilityResourceInner;

/** Samples for ResourceProvider CheckMigrationNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/CheckMigrationNameAvailability.json
     */
    /**
     * Sample code: CheckMigrationNameAvailability.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void checkMigrationNameAvailability(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .resourceProviders()
            .checkMigrationNameAvailabilityWithResponse(
                "ffffffff-ffff-ffff-ffff-ffffffffffff",
                "testrg",
                "testtarget",
                new MigrationNameAvailabilityResourceInner()
                    .withName("name1")
                    .withType("Microsoft.DBforPostgreSQL/flexibleServers/migrations"),
                com.azure.core.util.Context.NONE);
    }
}
