
/**
 * Samples for ServiceGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayList.json
     */
    /**
     * Sample code: List service gateway in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServiceGatewayInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
