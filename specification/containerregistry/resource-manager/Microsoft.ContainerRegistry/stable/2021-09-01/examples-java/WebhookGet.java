import com.azure.core.util.Context;

/** Samples for Webhooks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/WebhookGet.json
     */
    /**
     * Sample code: WebhookGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .getWithResponse("myResourceGroup", "myRegistry", "myWebhook", Context.NONE);
    }
}
