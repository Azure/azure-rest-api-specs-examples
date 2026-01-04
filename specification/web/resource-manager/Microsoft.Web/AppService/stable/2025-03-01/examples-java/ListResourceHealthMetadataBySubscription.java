
/**
 * Samples for ResourceHealthMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListResourceHealthMetadataBySubscription.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listResourceHealthMetadataForASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceHealthMetadatas().list(com.azure.core.util.Context.NONE);
    }
}
