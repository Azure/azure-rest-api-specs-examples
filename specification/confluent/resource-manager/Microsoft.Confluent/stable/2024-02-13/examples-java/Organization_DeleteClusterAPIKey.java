
/**
 * Samples for Organization DeleteClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_DeleteClusterAPIKey.json
     */
    /**
     * Sample code: Organization_DeleteClusterAPIKey.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationDeleteClusterAPIKey(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().deleteClusterApiKeyWithResponse("myResourceGroup", "myOrganization", "ZFZ6SZZZWGYBEIFB",
            com.azure.core.util.Context.NONE);
    }
}
