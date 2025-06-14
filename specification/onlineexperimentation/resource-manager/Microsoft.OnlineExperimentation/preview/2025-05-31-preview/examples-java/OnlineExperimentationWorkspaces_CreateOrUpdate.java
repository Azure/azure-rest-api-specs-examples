
import com.azure.resourcemanager.onlineexperimentation.models.ManagedServiceIdentity;
import com.azure.resourcemanager.onlineexperimentation.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.onlineexperimentation.models.OnlineExperimentationWorkspaceProperties;
import com.azure.resourcemanager.onlineexperimentation.models.OnlineExperimentationWorkspaceSku;
import com.azure.resourcemanager.onlineexperimentation.models.OnlineExperimentationWorkspaceSkuName;
import com.azure.resourcemanager.onlineexperimentation.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OnlineExperimentationWorkspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update an OnlineExperimentationWorkspace with Free sku.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void createOrUpdateAnOnlineExperimentationWorkspaceWithFreeSku(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.onlineExperimentationWorkspaces().define("expworkspace7").withRegion("eastus2")
            .withExistingResourceGroup("res9871").withTags(mapOf("newKey", "fakeTokenPlaceholder"))
            .withProperties(new OnlineExperimentationWorkspaceProperties().withLogAnalyticsWorkspaceResourceId(
                "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871")
                .withLogsExporterStorageAccountResourceId(
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871")
                .withAppConfigurationResourceId(
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                    new UserAssignedIdentity(),
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2",
                    new UserAssignedIdentity())))
            .withSku(new OnlineExperimentationWorkspaceSku().withName(OnlineExperimentationWorkspaceSkuName.F0))
            .create();
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
