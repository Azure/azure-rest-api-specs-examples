
/**
 * Samples for Webhooks GetCallbackConfig.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/WebhookGetCallbackConfig.json
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
