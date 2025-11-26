
import com.azure.resourcemanager.postgresqlflexibleserver.fluent.models.MigrationNameAvailabilityInner;

/**
 * Samples for Migrations CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * MigrationsCheckNameAvailability.json
     */
    /**
     * Sample code: Check the validity and availability of the given name, to assign it to a new migration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void checkTheValidityAndAvailabilityOfTheGivenNameToAssignItToANewMigration(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations()
            .checkNameAvailabilityWithResponse(
                "exampleresourcegroup", "exampleserver", new MigrationNameAvailabilityInner()
                    .withName("examplemigration").withType("Microsoft.DBforPostgreSQL/flexibleServers/migrations"),
                com.azure.core.util.Context.NONE);
    }
}
