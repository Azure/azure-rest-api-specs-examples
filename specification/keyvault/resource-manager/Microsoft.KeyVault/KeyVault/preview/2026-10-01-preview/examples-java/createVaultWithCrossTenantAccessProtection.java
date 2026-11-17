
import com.azure.resourcemanager.keyvault.models.Sku;
import com.azure.resourcemanager.keyvault.models.SkuFamily;
import com.azure.resourcemanager.keyvault.models.SkuName;
import com.azure.resourcemanager.keyvault.models.VaultCreateOrUpdateParameters;
import com.azure.resourcemanager.keyvault.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/createVaultWithCrossTenantAccessProtection.json
     */
    /**
     * Sample code: Create a vault with cross-tenant access protection enabled.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        createAVaultWithCrossTenantAccessProtectionEnabled(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().createOrUpdate("sample-resource-group", "sample-vault",
            new VaultCreateOrUpdateParameters().withLocation("westus").withProperties(
                new VaultProperties().withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                    .withSku(new Sku().withFamily(SkuFamily.A).withName(SkuName.STANDARD))
                    .withEnableRbacAuthorization(true).withEnableCrossTenantAccessProtection(true)),
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
