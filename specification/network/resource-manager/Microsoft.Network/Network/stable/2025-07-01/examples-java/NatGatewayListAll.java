
/**
 * Samples for NatGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayListAll.json
     */
    /**
     * Sample code: List all nat gateways.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNatGateways(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().list(com.azure.core.util.Context.NONE);
    }
}
