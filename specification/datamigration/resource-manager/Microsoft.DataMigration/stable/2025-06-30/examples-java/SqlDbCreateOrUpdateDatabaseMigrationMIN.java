
import com.azure.resourcemanager.datamigration.models.DatabaseMigrationPropertiesSqlDb;
import com.azure.resourcemanager.datamigration.models.SqlConnectionInformation;

/**
 * Samples for DatabaseMigrationsSqlDb CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * SqlDbCreateOrUpdateDatabaseMigrationMIN.json
     */
    /**
     * Sample code: Create or Update Database Migration resource with Minimum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateDatabaseMigrationResourceWithMinimumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlDbs().define("db1").withExistingServer("testrg", "sqldbinstance")
            .withProperties(new DatabaseMigrationPropertiesSqlDb().withScope(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance")
                .withMigrationService(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent")
                .withSourceSqlConnection(new SqlConnectionInformation().withDataSource("aaa")
                    .withAuthentication("WindowsAuthentication").withUsername("bbb")
                    .withPassword("fakeTokenPlaceholder").withEncryptConnection(true).withTrustServerCertificate(true))
                .withSourceDatabaseName("aaa")
                .withTargetSqlConnection(new SqlConnectionInformation().withDataSource("sqldbinstance")
                    .withAuthentication("SqlAuthentication").withUsername("bbb").withPassword("fakeTokenPlaceholder")
                    .withEncryptConnection(true).withTrustServerCertificate(true)))
            .create();
    }
}
