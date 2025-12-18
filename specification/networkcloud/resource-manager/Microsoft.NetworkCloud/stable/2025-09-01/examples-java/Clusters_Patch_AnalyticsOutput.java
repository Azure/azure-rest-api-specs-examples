
import com.azure.resourcemanager.networkcloud.models.AnalyticsOutputSettings;
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.IdentitySelector;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentity;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentitySelectorType;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.networkcloud.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_Patch_AnalyticsOutput.json
     */
    /**
     * Sample code: Patch cluster analytics output.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchClusterAnalyticsOutput(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
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
            .withAnalyticsOutputSettings(new AnalyticsOutputSettings().withAnalyticsWorkspaceId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName")
                .withAssociatedIdentity(new IdentitySelector()
                    .withIdentityType(ManagedServiceIdentitySelectorType.USER_ASSIGNED_IDENTITY)
                    .withUserAssignedIdentityResourceId(
                        "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2")))
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
