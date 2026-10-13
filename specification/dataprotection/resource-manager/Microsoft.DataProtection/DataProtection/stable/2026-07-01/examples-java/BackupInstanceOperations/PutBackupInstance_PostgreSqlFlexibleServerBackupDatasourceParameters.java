
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.BackupSolutionType;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.PolicyInfo;
import com.azure.resourcemanager.dataprotection.models.PolicyParameters;
import com.azure.resourcemanager.dataprotection.models.PostgreSqlFlexibleServerBackupDatasourceParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-07-01/BackupInstanceOperations/PutBackupInstance_PostgreSqlFlexibleServerBackupDatasourceParameters.json
     */
    /**
     * Sample code: Create BackupInstance With PostgreSqlFlexibleServerBackupDatasourceParameters.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithPostgreSqlFlexibleServerBackupDatasourceParameters(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("pgflexbi").withExistingBackupVault("pgflexrg", "pgflexvault")
            .withProperties(new BackupInstance().withFriendlyName("pgflexbi").withDataSourceInfo(new Datasource()
                .withDatasourceType("Microsoft.DBforPostgreSQL/flexibleServers").withObjectType("Datasource")
                .withResourceId(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver")
                .withResourceLocation("eastus2euap").withResourceName("pgflexserver")
                .withResourceType("Microsoft.DBforPostgreSQL/flexibleServers").withResourceUri(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.DBforPostgreSQL/flexibleServers").withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver")
                    .withResourceLocation("eastus2euap").withResourceName("pgflexserver")
                    .withResourceType("Microsoft.DBforPostgreSQL/flexibleServers").withResourceUri(
                        "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DataProtection/BackupVaults/pgflexvault/backupPolicies/pgflexpolicy")
                    .withPolicyParameters(new PolicyParameters().withBackupDatasourceParametersList(
                        Arrays.asList(new PostgreSqlFlexibleServerBackupDatasourceParameters()
                            .withBackupSolutionType(BackupSolutionType.PHYSICAL_BACKUP)))))
                .withObjectType("BackupInstance"))
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
