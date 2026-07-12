
import com.azure.resourcemanager.recoveryservices.models.ImmutabilityConfiguration;
import com.azure.resourcemanager.recoveryservices.models.ImmutabilitySettings;
import com.azure.resourcemanager.recoveryservices.models.ImmutabilityType;
import com.azure.resourcemanager.recoveryservices.models.SecuritySettings;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/PATCHVault_WithImmutabilityConfig.json
     */
    /**
     * Sample code: Update Vault With Immutability Config.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        updateVaultWithImmutabilityConfig(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource = manager.vaults()
            .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("PatchKey", "fakeTokenPlaceholder"))
            .withProperties(new VaultProperties().withSecuritySettings(
                new SecuritySettings().withImmutabilitySettings(new ImmutabilitySettings().withConfiguration(
                    new ImmutabilityConfiguration().withType(ImmutabilityType.TIME_BASED).withDurationInDays(30)))))
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
