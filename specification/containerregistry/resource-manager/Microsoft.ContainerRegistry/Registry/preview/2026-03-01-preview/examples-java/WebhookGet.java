
/**
 * Samples for Webhooks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/WebhookGet.json
     */
    /**
     * Sample code: WebhookGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void webhookGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().getWithResponse("myResourceGroup", "myRegistry", "myWebhook",
            com.azure.core.util.Context.NONE);
    }
}
