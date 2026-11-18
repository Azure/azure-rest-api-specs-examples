
import com.azure.resourcemanager.keyvault.models.VaultPatchParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/updateVaultPreservingIdentity.json
     */
    /**
     * Sample code: Preserve existing user-assigned identities by omitting identity from PATCH.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void preserveExistingUserAssignedIdentitiesByOmittingIdentityFromPATCH(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().updateWithResponse("sample-resource-group", "sample-vault",
            new VaultPatchParameters().withTags(mapOf("environment", "production")), com.azure.core.util.Context.NONE);
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
