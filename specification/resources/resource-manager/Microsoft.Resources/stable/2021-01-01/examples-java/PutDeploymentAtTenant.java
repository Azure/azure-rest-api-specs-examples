import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.ScopedDeployment;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Deployments CreateOrUpdateAtTenantScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PutDeploymentAtTenant.json
     */
    /**
     * Sample code: Create deployment at tenant scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDeploymentAtTenantScope(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getDeployments()
            .createOrUpdateAtTenantScope(
                "tenant-dep01",
                new ScopedDeployment()
                    .withLocation("eastus")
                    .withProperties(
                        new DeploymentProperties()
                            .withTemplateLink(new TemplateLink().withUri("https://example.com/exampleTemplate.json"))
                            .withParameters(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                            .withMode(DeploymentMode.INCREMENTAL))
                    .withTags(mapOf("tagKey1", "tag-value-1", "tagKey2", "tag-value-2")),
                Context.NONE);
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
