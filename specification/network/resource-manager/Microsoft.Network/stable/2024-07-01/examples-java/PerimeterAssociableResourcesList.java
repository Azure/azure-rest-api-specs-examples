
/**
 * Samples for NetworkSecurityPerimeterAssociableResourceTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * PerimeterAssociableResourcesList.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterAssociableResourceTypes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        networkSecurityPerimeterAssociableResourceTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAssociableResourceTypes().list("westus",
            com.azure.core.util.Context.NONE);
    }
}
