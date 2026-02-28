
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
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookGetCallbackConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().getCallbackConfigWithResponse(
            "myResourceGroup", "myRegistry", "myWebhook", com.azure.core.util.Context.NONE);
    }
}
