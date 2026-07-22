
/**
 * Samples for Webhooks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/WebhookList.json
     */
    /**
     * Sample code: WebhookList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void webhookList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().list("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
