
import com.azure.resourcemanager.keyvault.models.ManagedServiceIdentity;
import com.azure.resourcemanager.keyvault.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.keyvault.models.VaultPatchParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/removeVaultWithIdentity.json
     */
    /**
     * Sample code: Remove all user-assigned identities and their certificates from a vault using PATCH.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void removeAllUserAssignedIdentitiesAndTheirCertificatesFromAVaultUsingPATCH(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().updateWithResponse("sample-resource-group", "sample-vault",
            new VaultPatchParameters()
                .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)),
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
