
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
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void webhookDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().delete("myResourceGroup", "myRegistry", "myWebhook",
            com.azure.core.util.Context.NONE);
    }
}
