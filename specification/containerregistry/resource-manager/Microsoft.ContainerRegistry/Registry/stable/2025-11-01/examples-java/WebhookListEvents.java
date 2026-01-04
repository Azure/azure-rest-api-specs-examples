
/**
 * Samples for Webhooks ListEvents.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * WebhookListEvents.json
     */
    /**
     * Sample code: WebhookListEvents.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookListEvents(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().listEvents("myResourceGroup", "myRegistry",
            "myWebhook", com.azure.core.util.Context.NONE);
    }
}
