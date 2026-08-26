
import com.azure.resourcemanager.recoveryservices.models.RegionOfChoiceSettings;
import com.azure.resourcemanager.recoveryservices.models.State;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/PATCHVault_WithRegionOfChoiceSettings.json
     */
    /**
     * Sample code: Update Vault With Region Of Choice Settings.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateVaultWithRegionOfChoiceSettings(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource = manager.vaults()
            .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("PatchKey", "fakeTokenPlaceholder")).withProperties(
            new VaultProperties().withRegionOfChoiceSettings(new RegionOfChoiceSettings().withStatus(State.ENABLED)))
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
