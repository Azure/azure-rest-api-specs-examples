
/**
 * Samples for Webhooks ListEvents.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/WebhookListEvents.json
     */
    /**
     * Sample code: WebhookListEvents.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void webhookListEvents(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().listEvents("myResourceGroup", "myRegistry", "myWebhook",
            com.azure.core.util.Context.NONE);
    }
}
