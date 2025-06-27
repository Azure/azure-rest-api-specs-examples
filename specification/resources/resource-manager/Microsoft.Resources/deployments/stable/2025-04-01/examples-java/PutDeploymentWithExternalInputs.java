
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.DeploymentInner;
import com.azure.resourcemanager.resources.models.DeploymentExternalInput;
import com.azure.resourcemanager.resources.models.DeploymentExternalInputDefinition;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentParameter;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/
     * PutDeploymentWithExternalInputs.json
     */
    /**
     * Sample code: Create deployment using external inputs.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDeploymentUsingExternalInputs(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure.genericResources().manager().deploymentClient().getDeployments().createOrUpdate("my-resource-group",
            "my-deployment",
            new DeploymentInner().withProperties(new DeploymentProperties()
                .withTemplate(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"$schema\":\"https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#\",\"contentVersion\":\"1.0.0.0\",\"outputs\":{\"inputObj\":{\"type\":\"object\",\"value\":\"[parameters('inputObj')]\"}},\"parameters\":{\"inputObj\":{\"type\":\"object\"}},\"resources\":[]}",
                    Object.class, SerializerEncoding.JSON))
                .withParameters(mapOf("inputObj",
                    new DeploymentParameter().withExpression("[createObject('foo', externalInputs('fooValue'))]")))
                .withExternalInputs(mapOf("fooValue", new DeploymentExternalInput().withValue("baz")))
                .withExternalInputDefinitions(mapOf("fooValue",
                    new DeploymentExternalInputDefinition().withKind("sys.envVar").withConfig("FOO_VALUE")))
                .withMode(DeploymentMode.INCREMENTAL)),
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
