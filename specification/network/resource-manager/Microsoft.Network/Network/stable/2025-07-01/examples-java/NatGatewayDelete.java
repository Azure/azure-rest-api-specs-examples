
/**
 * Samples for NatGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayDelete.json
     */
    /**
     * Sample code: Delete nat gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNatGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().delete("rg1", "test-natGateway", com.azure.core.util.Context.NONE);
    }
}
