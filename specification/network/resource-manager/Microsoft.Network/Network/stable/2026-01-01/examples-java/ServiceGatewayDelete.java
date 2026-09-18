
/**
 * Samples for ServiceGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ServiceGatewayDelete.json
     */
    /**
     * Sample code: Delete service gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteServiceGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().delete("rg1", "sg", com.azure.core.util.Context.NONE);
    }
}
