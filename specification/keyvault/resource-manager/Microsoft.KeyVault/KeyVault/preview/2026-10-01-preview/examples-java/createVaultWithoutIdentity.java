
import com.azure.resourcemanager.keyvault.models.ManagedServiceIdentity;
import com.azure.resourcemanager.keyvault.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.keyvault.models.Sku;
import com.azure.resourcemanager.keyvault.models.SkuFamily;
import com.azure.resourcemanager.keyvault.models.SkuName;
import com.azure.resourcemanager.keyvault.models.VaultCreateOrUpdateParameters;
import com.azure.resourcemanager.keyvault.models.VaultProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/createVaultWithoutIdentity.json
     */
    /**
     * Sample code: Remove all user-assigned identities and their certificates from a vault using PUT.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void removeAllUserAssignedIdentitiesAndTheirCertificatesFromAVaultUsingPUT(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().createOrUpdate("sample-resource-group", "sample-vault",
            new VaultCreateOrUpdateParameters()
                .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE))
                .withLocation("westus").withProperties(
                    new VaultProperties().withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                        .withSku(new Sku().withFamily(SkuFamily.A).withName(SkuName.STANDARD))
                        .withAccessPolicies(Arrays.asList()).withEnableRbacAuthorization(true)),
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
