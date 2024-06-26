import com.azure.core.util.Context;

/** Samples for ApplicationGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ApplicationGatewayDelete.json
     */
    /**
     * Sample code: Delete ApplicationGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().delete("rg1", "appgw", Context.NONE);
    }
}
