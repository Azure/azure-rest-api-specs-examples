import com.azure.core.util.Context;

/** Samples for Webhooks ListEvents. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/WebhookListEvents.json
     */
    /**
     * Sample code: WebhookListEvents.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookListEvents(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .listEvents("myResourceGroup", "myRegistry", "myWebhook", Context.NONE);
    }
}
