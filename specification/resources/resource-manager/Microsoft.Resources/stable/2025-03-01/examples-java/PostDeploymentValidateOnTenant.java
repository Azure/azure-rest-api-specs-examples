
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.ScopedDeployment;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Deployments ValidateAtTenantScope.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2025-03-01/examples/
     * PostDeploymentValidateOnTenant.json
     */
    /**
     * Sample code: Validates a template at tenant scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validatesATemplateAtTenantScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getDeployments().validateAtTenantScope("my-deployment",
            new ScopedDeployment().withLocation("eastus")
                .withProperties(new DeploymentProperties()
                    .withTemplateLink(new TemplateLink().withUri("https://example.com/exampleTemplate.json"))
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
