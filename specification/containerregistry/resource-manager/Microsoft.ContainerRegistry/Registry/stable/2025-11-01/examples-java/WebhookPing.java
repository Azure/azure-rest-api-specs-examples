
/**
 * Samples for Webhooks Ping.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * WebhookPing.json
     */
    /**
     * Sample code: WebhookPing.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookPing(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().pingWithResponse("myResourceGroup",
            "myRegistry", "myWebhook", com.azure.core.util.Context.NONE);
    }
}
