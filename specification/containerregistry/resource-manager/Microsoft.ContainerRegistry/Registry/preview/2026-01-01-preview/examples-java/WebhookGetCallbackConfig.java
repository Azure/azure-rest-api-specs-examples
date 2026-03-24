
/**
 * Samples for Webhooks GetCallbackConfig.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/WebhookGetCallbackConfig.json
     */
    /**
     * Sample code: WebhookGetCallbackConfig.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        webhookGetCallbackConfig(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getWebhooks().getCallbackConfigWithResponse("myResourceGroup", "myRegistry",
            "myWebhook", com.azure.core.util.Context.NONE);
    }
}
