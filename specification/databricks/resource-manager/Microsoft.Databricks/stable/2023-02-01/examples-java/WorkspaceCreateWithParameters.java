import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomStringParameter;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceCreateWithParameters.json
     */
    /**
     * Sample code: Create or update workspace with custom parameters.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createOrUpdateWorkspaceWithCustomParameters(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withParameters(
                new WorkspaceCustomParameters()
                    .withCustomVirtualNetworkId(
                        new WorkspaceCustomStringParameter()
                            .withValue(
                                "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myNetwork"))
                    .withCustomPublicSubnetName(new WorkspaceCustomStringParameter().withValue("myPublicSubnet"))
                    .withCustomPrivateSubnetName(new WorkspaceCustomStringParameter().withValue("myPrivateSubnet")))
            .create();
    }

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
