
/** Samples for Queue Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/QueueOperationGet.json
     */
    /**
     * Sample code: QueueOperationGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getQueues().getWithResponse("res3376", "sto328", "queue6185",
            com.azure.core.util.Context.NONE);
    }
}
