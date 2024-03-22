
/**
 * Samples for Organization GetClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_GetClusterAPIKey.json
     */
    /**
     * Sample code: Organization_GetClusterAPIKey.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetClusterAPIKey(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getClusterApiKeyWithResponse("myResourceGroup", "myOrganization", "apiKeyId-123",
            com.azure.core.util.Context.NONE);
    }
}
