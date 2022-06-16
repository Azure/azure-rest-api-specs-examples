import com.azure.core.util.Context;

/** Samples for Webhooks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/WebhookDelete.json
     */
    /**
     * Sample code: WebhookDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .delete("myResourceGroup", "myRegistry", "myWebhook", Context.NONE);
    }
}
