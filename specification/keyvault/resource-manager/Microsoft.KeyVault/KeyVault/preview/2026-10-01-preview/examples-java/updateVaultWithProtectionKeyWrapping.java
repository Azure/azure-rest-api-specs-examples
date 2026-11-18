
import com.azure.resourcemanager.keyvault.models.VaultPatchParameters;
import com.azure.resourcemanager.keyvault.models.VaultPatchProperties;
import com.azure.resourcemanager.keyvault.models.WrappingKeyAuthenticationData;
import com.azure.resourcemanager.keyvault.models.WrappingKeyConfiguration;
import com.azure.resourcemanager.keyvault.models.WrappingKeyValue;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/updateVaultWithProtectionKeyWrapping.json
     */
    /**
     * Sample code: Configure protection key wrapping on an existing vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        configureProtectionKeyWrappingOnAnExistingVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().updateWithResponse("sample-resource-group", "sample-vault",
            new VaultPatchParameters()
                .withProperties(new VaultPatchProperties().withProtectionKeysWrappingKey(new WrappingKeyConfiguration()
                    .withKeyValue(
                        new WrappingKeyValue().withUrl("https://sample-hsm.managedhsm.azure.net/keys/wrapping-key"))
                    .withKeyAuthenticationData(new WrappingKeyAuthenticationData()))),
            com.azure.core.util.Context.NONE);
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
