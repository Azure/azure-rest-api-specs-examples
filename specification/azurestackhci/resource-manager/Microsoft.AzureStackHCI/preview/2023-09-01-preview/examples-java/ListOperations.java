/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listOperations(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
