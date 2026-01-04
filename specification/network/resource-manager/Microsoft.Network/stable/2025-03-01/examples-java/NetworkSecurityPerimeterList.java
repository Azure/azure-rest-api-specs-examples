
/**
 * Samples for NetworkSecurityPerimeters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkSecurityPerimeterList.
     * json
     */
    /**
     * Sample code: List Network Security Perimeter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkSecurityPerimeter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeters().listByResourceGroup("rg1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
