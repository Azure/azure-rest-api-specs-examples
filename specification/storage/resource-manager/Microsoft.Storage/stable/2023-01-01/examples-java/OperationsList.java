
/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void operationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
