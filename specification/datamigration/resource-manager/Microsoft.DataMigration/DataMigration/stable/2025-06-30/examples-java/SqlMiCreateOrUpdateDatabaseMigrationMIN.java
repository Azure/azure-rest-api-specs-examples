
import com.azure.resourcemanager.datamigration.models.BackupConfiguration;
import com.azure.resourcemanager.datamigration.models.DatabaseMigrationPropertiesSqlMi;
import com.azure.resourcemanager.datamigration.models.SourceLocation;
import com.azure.resourcemanager.datamigration.models.SqlConnectionInformation;
import com.azure.resourcemanager.datamigration.models.SqlFileShare;
import com.azure.resourcemanager.datamigration.models.TargetLocation;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DatabaseMigrationsSqlMi CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * SqlMiCreateOrUpdateDatabaseMigrationMIN.json
     */
    /**
     * Sample code: Create or Update Database Migration resource with Minimum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateDatabaseMigrationResourceWithMinimumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlMis().define("db1").withExistingManagedInstance("testrg", "managedInstance1")
            .withProperties(new DatabaseMigrationPropertiesSqlMi().withScope(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance")
                .withMigrationService(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent")
                .withSourceSqlConnection(new SqlConnectionInformation().withDataSource("aaa")
                    .withAuthentication("WindowsAuthentication").withUsername("bbb")
                    .withPassword("fakeTokenPlaceholder").withEncryptConnection(true).withTrustServerCertificate(true))
                .withSourceDatabaseName("aaa")
                .withBackupConfiguration(new BackupConfiguration()
                    .withSourceLocation(new SourceLocation().withFileShare(new SqlFileShare()
                        .withPath("C:\\aaa\\bbb\\ccc").withUsername("name").withPassword("fakeTokenPlaceholder")))
                    .withTargetLocation(
                        new TargetLocation().withStorageAccountResourceId("account.database.windows.net")
                            .withAccountKey("fakeTokenPlaceholder"))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
