
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.CommandOutputOverride;
import com.azure.resourcemanager.networkcloud.models.CommandOutputSettings;
import com.azure.resourcemanager.networkcloud.models.CommandOutputType;
import com.azure.resourcemanager.networkcloud.models.IdentitySelector;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentity;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentitySelectorType;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.networkcloud.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_Patch_CommandOutput.json
     */
    /**
     * Sample code: Patch cluster command output.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchClusterCommandOutput(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1",
                    null,
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2",
                    new UserAssignedIdentity())))
            .withCommandOutputSettings(new CommandOutputSettings().withAssociatedIdentity(new IdentitySelector()
                .withIdentityType(ManagedServiceIdentitySelectorType.USER_ASSIGNED_IDENTITY)
                .withUserAssignedIdentityResourceId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2"))
                .withContainerUrl("https://myaccount.blob.core.windows.net/mycontainer?restype=container")
                .withOverrides(Arrays.asList(new CommandOutputOverride().withAssociatedIdentity(new IdentitySelector()
                    .withIdentityType(ManagedServiceIdentitySelectorType.USER_ASSIGNED_IDENTITY)
                    .withUserAssignedIdentityResourceId(
                        "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2"))
                    .withCommandOutputType(CommandOutputType.STORAGE_RUN_READ_COMMANDS)
                    .withContainerUrl("https://myaccount.blob.core.windows.net/myContainer2?restype=container"))))
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
