
/**
 * Samples for Webhooks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * WebhookGet.json
     */
    /**
     * Sample code: WebhookGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getWebhooks().getWithResponse("myResourceGroup",
            "myRegistry", "myWebhook", com.azure.core.util.Context.NONE);
    }
}
