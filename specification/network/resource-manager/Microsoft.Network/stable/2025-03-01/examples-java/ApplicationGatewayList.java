
/**
 * Samples for ApplicationGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationGatewayList.json
     */
    /**
     * Sample code: Lists all application gateways in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listsAllApplicationGatewaysInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
