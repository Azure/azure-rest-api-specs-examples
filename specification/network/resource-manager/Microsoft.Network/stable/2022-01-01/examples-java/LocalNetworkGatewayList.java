import com.azure.core.util.Context;

/** Samples for LocalNetworkGateways ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/LocalNetworkGatewayList.json
     */
    /**
     * Sample code: ListLocalNetworkGateways.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLocalNetworkGateways(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLocalNetworkGateways().listByResourceGroup("rg1", Context.NONE);
    }
}
