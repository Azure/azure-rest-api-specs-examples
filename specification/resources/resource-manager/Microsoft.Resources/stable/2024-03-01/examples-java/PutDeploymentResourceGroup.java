
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
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * PutDeploymentResourceGroup.json
     */
    /**
     * Sample code: Create a deployment that will deploy a template with a uri and queryString.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADeploymentThatWillDeployATemplateWithAUriAndQueryString(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getDeployments().createOrUpdate("my-resource-group",
            "my-deployment",
            new DeploymentInner().withProperties(new DeploymentProperties().withTemplateLink(
                new TemplateLink().withUri("https://example.com/exampleTemplate.json").withQueryString(
                    "sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=xxxxxxxx0xxxxxxxxxxxxx%2bxxxxxxxxxxxxxxxxxxxx%3d"))
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
