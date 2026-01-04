
/**
 * Samples for NetworkSecurityPerimeters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkSecurityPerimeterGet.
     * json
     */
    /**
     * Sample code: NetworkSecurityPerimeterGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkSecurityPerimeterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeters().getByResourceGroupWithResponse("rg1",
            "nsp1", com.azure.core.util.Context.NONE);
    }
}
