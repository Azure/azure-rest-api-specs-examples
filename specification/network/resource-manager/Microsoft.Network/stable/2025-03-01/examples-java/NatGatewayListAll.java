
/**
 * Samples for NatGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NatGatewayListAll.json
     */
    /**
     * Sample code: List all nat gateways.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNatGateways(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatGateways().list(com.azure.core.util.Context.NONE);
    }
}
