import com.azure.core.util.Context;

/** Samples for Webhooks GetCallbackConfig. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/WebhookGetCallbackConfig.json
     */
    /**
     * Sample code: WebhookGetCallbackConfig.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookGetCallbackConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .getCallbackConfigWithResponse("myResourceGroup", "myRegistry", "myWebhook", Context.NONE);
    }
}
