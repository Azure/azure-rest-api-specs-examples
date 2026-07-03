
/**
 * Samples for ServiceGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceGatewayList.json
     */
    /**
     * Sample code: List service gateway in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listServiceGatewayInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
