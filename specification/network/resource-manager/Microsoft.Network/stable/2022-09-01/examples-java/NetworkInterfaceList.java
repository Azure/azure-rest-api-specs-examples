/** Samples for NetworkInterfaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkInterfaceList.json
     */
    /**
     * Sample code: List network interfaces in resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkInterfacesInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
