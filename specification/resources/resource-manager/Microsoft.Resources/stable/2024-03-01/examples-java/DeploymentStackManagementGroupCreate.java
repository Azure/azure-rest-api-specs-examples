
import com.azure.resourcemanager.resources.fluent.models.DeploymentStackInner;
import com.azure.resourcemanager.resources.models.ActionOnUnmanage;
import com.azure.resourcemanager.resources.models.DenySettings;
import com.azure.resourcemanager.resources.models.DenySettingsMode;
import com.azure.resourcemanager.resources.models.DeploymentParameter;
import com.azure.resourcemanager.resources.models.DeploymentStacksDeleteDetachEnum;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DeploymentStacks CreateOrUpdateAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackManagementGroupCreate.json
     */
    /**
     * Sample code: DeploymentStacksManagementGroupCreateOrUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deploymentStacksManagementGroupCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .createOrUpdateAtManagementGroup("myMg", "simpleDeploymentStack",
                new DeploymentStackInner().withLocation("eastus").withTags(mapOf("tagkey", "fakeTokenPlaceholder"))
                    .withParameters(mapOf("parameter1", new DeploymentParameter().withValue("a string")))
                    .withActionOnUnmanage(new ActionOnUnmanage().withResources(DeploymentStacksDeleteDetachEnum.DELETE)
                        .withResourceGroups(DeploymentStacksDeleteDetachEnum.DELETE)
                        .withManagementGroups(DeploymentStacksDeleteDetachEnum.DETACH))
                    .withDenySettings(new DenySettings().withMode(DenySettingsMode.DENY_DELETE)
                        .withExcludedPrincipals(Arrays.asList("principal")).withExcludedActions(Arrays.asList("action"))
                        .withApplyToChildScopes(false)),
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
