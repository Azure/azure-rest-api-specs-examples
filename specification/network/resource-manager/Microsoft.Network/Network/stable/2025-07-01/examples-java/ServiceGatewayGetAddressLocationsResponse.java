
/**
 * Samples for ServiceGateways GetAddressLocations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceGatewayGetAddressLocationsResponse.json
     */
    /**
     * Sample code: Get address locations in service gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAddressLocationsInServiceGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().getAddressLocations("rg1", "sg", com.azure.core.util.Context.NONE);
    }
}
