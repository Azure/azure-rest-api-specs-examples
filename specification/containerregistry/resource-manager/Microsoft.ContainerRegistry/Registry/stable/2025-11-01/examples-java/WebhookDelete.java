
/**
 * Samples for Webhooks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/WebhookDelete.json
     */
    /**
     * Sample code: WebhookDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().delete("myResourceGroup", "myRegistry",
            "myWebhook", com.azure.core.util.Context.NONE);
    }
}
