
import com.azure.resourcemanager.datamigration.models.AuthType;
import com.azure.resourcemanager.datamigration.models.AzureBlob;
import com.azure.resourcemanager.datamigration.models.BackupConfiguration;
import com.azure.resourcemanager.datamigration.models.DatabaseMigrationPropertiesSqlMi;
import com.azure.resourcemanager.datamigration.models.ManagedServiceIdentity;
import com.azure.resourcemanager.datamigration.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.datamigration.models.OfflineConfiguration;
import com.azure.resourcemanager.datamigration.models.SourceLocation;
import com.azure.resourcemanager.datamigration.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DatabaseMigrationsSqlMi CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * SqlMiCreateOrUpdateDatabaseMigrationBlobManagedIdentity.json
     */
    /**
     * Sample code: Create or Update Database Migration resource from Azure Blob using Managed Identity.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateDatabaseMigrationResourceFromAzureBlobUsingManagedIdentity(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlMis().define("db1").withExistingManagedInstance("testrg", "managedInstance1")
            .withProperties(new DatabaseMigrationPropertiesSqlMi().withScope(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance")
                .withMigrationService(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent")
                .withSourceDatabaseName("aaa")
                .withBackupConfiguration(new BackupConfiguration().withSourceLocation(
                    new SourceLocation().withAzureBlob(new AzureBlob().withAuthType(AuthType.MANAGED_IDENTITY)
                        .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentities(mapOf(
                                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity",
                                new UserAssignedIdentity())))
                        .withStorageAccountResourceId(
                            "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                        .withBlobContainerName("test"))))
                .withOfflineConfiguration(
                    new OfflineConfiguration().withOffline(true).withLastBackupName("last_backup_file_name")))
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
