
import com.azure.resourcemanager.recoveryservices.models.CrossRegionRestore;
import com.azure.resourcemanager.recoveryservices.models.StandardTierStorageRedundancy;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import com.azure.resourcemanager.recoveryservices.models.VaultPropertiesRedundancySettings;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * PATCHVault_WithRedundancySettings.json
     */
    /**
     * Sample code: Update Vault With Redundancy Setting.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        updateVaultWithRedundancySetting(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource = manager.vaults()
            .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new VaultProperties().withRedundancySettings(new VaultPropertiesRedundancySettings()
                .withStandardTierStorageRedundancy(StandardTierStorageRedundancy.GEO_REDUNDANT)
                .withCrossRegionRestore(CrossRegionRestore.ENABLED)))
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
