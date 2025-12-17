
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.IdentitySelector;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentitySelectorType;
import com.azure.resourcemanager.networkcloud.models.SecretArchiveSettings;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_Patch_SecretArchive.json
     */
    /**
     * Sample code: Patch secret archive.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchSecretArchive(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withSecretArchiveSettings(new SecretArchiveSettings().withAssociatedIdentity(new IdentitySelector()
                .withIdentityType(ManagedServiceIdentitySelectorType.USER_ASSIGNED_IDENTITY)
                .withUserAssignedIdentityResourceId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1"))
                .withVaultUri("https://keyvaultname.vault.azure.net/"))
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
