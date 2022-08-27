import com.azure.core.util.Context;

/** Samples for QueueServices GetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/QueueServicesGet.json
     */
    /**
     * Sample code: QueueServicesGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueServicesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueueServices()
            .getServicePropertiesWithResponse("res4410", "sto8607", Context.NONE);
    }
}
