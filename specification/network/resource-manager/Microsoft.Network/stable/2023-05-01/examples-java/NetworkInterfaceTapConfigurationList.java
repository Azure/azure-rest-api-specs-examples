/** Samples for NetworkInterfaceTapConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkInterfaceTapConfigurationList.json
     */
    /**
     * Sample code: List virtual network tap configurations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworkTapConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaceTapConfigurations()
            .list("rg1", "mynic", com.azure.core.util.Context.NONE);
    }
}
