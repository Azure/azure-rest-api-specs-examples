
/**
 * Samples for NatGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NatGatewayList.json
     */
    /**
     * Sample code: List nat gateways in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNatGatewaysInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
