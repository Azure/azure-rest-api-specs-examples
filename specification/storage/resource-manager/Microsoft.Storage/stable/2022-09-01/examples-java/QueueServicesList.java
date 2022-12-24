import com.azure.core.util.Context;

/** Samples for QueueServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/QueueServicesList.json
     */
    /**
     * Sample code: QueueServicesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueueServices()
            .listWithResponse("res9290", "sto1590", Context.NONE);
    }
}
