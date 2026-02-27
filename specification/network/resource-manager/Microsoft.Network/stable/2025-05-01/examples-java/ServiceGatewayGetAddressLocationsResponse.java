
/**
 * Samples for ServiceGateways GetAddressLocations.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ServiceGatewayGetAddressLocationsResponse.json
     */
    /**
     * Sample code: Get address locations in service gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAddressLocationsInServiceGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceGateways().getAddressLocations("rg1", "sg",
            com.azure.core.util.Context.NONE);
    }
}
