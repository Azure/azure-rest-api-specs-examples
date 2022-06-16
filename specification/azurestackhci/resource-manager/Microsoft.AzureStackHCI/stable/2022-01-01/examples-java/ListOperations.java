import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-01-01/examples/ListOperations.json
     */
    /**
     * Sample code: List operations available with the Microsoft.AzureStackHCI provider.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listOperationsAvailableWithTheMicrosoftAzureStackHCIProvider(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
