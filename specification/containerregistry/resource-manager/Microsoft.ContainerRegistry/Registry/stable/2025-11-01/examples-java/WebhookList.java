
/**
 * Samples for Webhooks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/WebhookList.json
     */
    /**
     * Sample code: WebhookList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
