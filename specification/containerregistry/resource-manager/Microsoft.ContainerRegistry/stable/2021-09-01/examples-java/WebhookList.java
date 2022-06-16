import com.azure.core.util.Context;

/** Samples for Webhooks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/WebhookList.json
     */
    /**
     * Sample code: WebhookList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}
