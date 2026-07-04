
/**
 * Samples for ServiceGateways GetServices.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceGatewayGetServicesResponse.json
     */
    /**
     * Sample code: Get services in service gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getServicesInServiceGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().getServices("rg1", "sg", com.azure.core.util.Context.NONE);
    }
}
