/** Samples for NetworkInterfacesOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListNetworkInterfaceByResourceGroup.json
     */
    /**
     * Sample code: ListNetworkInterfaceByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listNetworkInterfaceByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.networkInterfacesOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
