
/**
 * Samples for NatGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayList.json
     */
    /**
     * Sample code: List nat gateways in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNatGatewaysInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
