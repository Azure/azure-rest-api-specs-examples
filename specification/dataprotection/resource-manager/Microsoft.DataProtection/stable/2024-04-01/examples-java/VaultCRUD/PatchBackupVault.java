
import com.azure.resourcemanager.dataprotection.models.AlertsState;
import com.azure.resourcemanager.dataprotection.models.AzureMonitorAlertSettings;
import com.azure.resourcemanager.dataprotection.models.BackupVaultResource;
import com.azure.resourcemanager.dataprotection.models.MonitoringSettings;
import com.azure.resourcemanager.dataprotection.models.PatchBackupVaultInput;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupVaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/
     * PatchBackupVault.json
     */
    /**
     * Sample code: Patch BackupVault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void patchBackupVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        BackupVaultResource resource = manager.backupVaults()
            .getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("newKey", "fakeTokenPlaceholder"))
            .withProperties(new PatchBackupVaultInput()
                .withMonitoringSettings(new MonitoringSettings().withAzureMonitorAlertSettings(
                    new AzureMonitorAlertSettings().withAlertsForAllJobFailures(AlertsState.ENABLED))))
            .apply();
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
