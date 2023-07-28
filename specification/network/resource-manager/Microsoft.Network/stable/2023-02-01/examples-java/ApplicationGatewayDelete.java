/** Samples for ApplicationGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayDelete.json
     */
    /**
     * Sample code: Delete ApplicationGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .delete("rg1", "appgw", com.azure.core.util.Context.NONE);
    }
}
