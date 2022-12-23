import com.azure.core.util.Context;

/** Samples for Queue List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/QueueOperationList.json
     */
    /**
     * Sample code: QueueOperationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueues()
            .list("res9290", "sto328", null, null, Context.NONE);
    }
}
