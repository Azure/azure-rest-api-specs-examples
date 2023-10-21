/** Samples for NetworkInterfacesOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetNetworkInterface.json
     */
    /**
     * Sample code: GetNetworkInterface.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getNetworkInterface(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .networkInterfacesOperations()
            .getByResourceGroupWithResponse("test-rg", "test-nic", com.azure.core.util.Context.NONE);
    }
}
