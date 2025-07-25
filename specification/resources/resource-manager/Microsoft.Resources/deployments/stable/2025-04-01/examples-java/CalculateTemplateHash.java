
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for Deployments CalculateTemplateHash.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/
     * CalculateTemplateHash.json
     */
    /**
     * Sample code: Calculate template hash.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void calculateTemplateHash(com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure.genericResources().manager().deploymentClient().getDeployments()
            .calculateTemplateHashWithResponse(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                "{\"$schema\":\"http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview\",\"contentVersion\":\"1.0.0.0\",\"outputs\":{\"string\":{\"type\":\"string\",\"value\":\"myvalue\"}},\"parameters\":{\"string\":{\"type\":\"string\"}},\"resources\":[],\"variables\":{\"array\":[1,2,3,4],\"bool\":true,\"int\":42,\"object\":{\"object\":{\"location\":\"West US\",\"vmSize\":\"Large\"}},\"string\":\"string\"}}",
                Object.class, SerializerEncoding.JSON), com.azure.core.util.Context.NONE);
    }
}
