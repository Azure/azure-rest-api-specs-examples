/** Samples for NetworkInterfacesOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteNetworkInterface.json
     */
    /**
     * Sample code: DeleteNetworkInterface.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteNetworkInterface(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.networkInterfacesOperations().delete("test-rg", "test-nic", com.azure.core.util.Context.NONE);
    }
}
