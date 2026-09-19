
import com.azure.resourcemanager.dataprotection.models.BackupVaultResource;
import com.azure.resourcemanager.dataprotection.models.CostManagementSettings;
import com.azure.resourcemanager.dataprotection.models.GranularityLevel;
import com.azure.resourcemanager.dataprotection.models.PatchBackupVaultInput;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupVaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/VaultCRUD/PatchBackupVaultWithCostManagementSettings.json
     */
    /**
     * Sample code: Patch BackupVault with Cost Management Settings.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void patchBackupVaultWithCostManagementSettings(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        BackupVaultResource resource = manager.backupVaults()
            .getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("newKey", "fakeTokenPlaceholder"))
            .withProperties(new PatchBackupVaultInput().withCostManagementSettings(
                new CostManagementSettings().withGranularityLevel(GranularityLevel.PROTECTED_ITEM_LEVEL)))
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
