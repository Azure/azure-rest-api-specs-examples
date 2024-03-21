
import com.azure.resourcemanager.confluent.models.CreateApiKeyModel;

/**
 * Samples for Organization CreateApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_CreateClusterAPIKey.json
     */
    /**
     * Sample code: Organization_CreateAPIKey.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationCreateAPIKey(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().createApiKeyWithResponse(
            "myResourceGroup", "myOrganization", "env-12132", "clusterId-123", new CreateApiKeyModel()
                .withName("CI kafka access key").withDescription("This API key provides kafka access to cluster x"),
            com.azure.core.util.Context.NONE);
    }
}
