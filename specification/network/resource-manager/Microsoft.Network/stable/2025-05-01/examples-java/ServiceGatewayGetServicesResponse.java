
/**
 * Samples for ServiceGateways GetServices.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ServiceGatewayGetServicesResponse.json
     */
    /**
     * Sample code: Get services in service gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServicesInServiceGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceGateways().getServices("rg1", "sg",
            com.azure.core.util.Context.NONE);
    }
}
