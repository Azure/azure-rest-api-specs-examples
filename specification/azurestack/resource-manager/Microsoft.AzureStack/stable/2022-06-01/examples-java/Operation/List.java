/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Operation/List.json
     */
    /**
     * Sample code: Returns the list of supported REST operations.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsTheListOfSupportedRESTOperations(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
