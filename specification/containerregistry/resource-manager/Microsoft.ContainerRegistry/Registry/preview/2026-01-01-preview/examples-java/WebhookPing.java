
/**
 * Samples for Webhooks Ping.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/WebhookPing.json
     */
    /**
     * Sample code: WebhookPing.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void webhookPing(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().pingWithResponse("myResourceGroup", "myRegistry", "myWebhook",
            com.azure.core.util.Context.NONE);
    }
}
