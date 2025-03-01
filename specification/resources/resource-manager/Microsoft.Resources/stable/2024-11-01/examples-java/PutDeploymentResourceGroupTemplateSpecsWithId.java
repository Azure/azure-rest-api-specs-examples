
import com.azure.resourcemanager.resources.fluent.models.DeploymentInner;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-11-01/examples/
     * PutDeploymentResourceGroupTemplateSpecsWithId.json
     */
    /**
     * Sample code: Create a deployment that will deploy a templateSpec with the given resourceId.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADeploymentThatWillDeployATemplateSpecWithTheGivenResourceId(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getDeployments().createOrUpdate("my-resource-group",
            "my-deployment",
            new DeploymentInner().withProperties(new DeploymentProperties().withTemplateLink(new TemplateLink().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"))
                .withParameters(mapOf()).withMode(DeploymentMode.INCREMENTAL)),
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
