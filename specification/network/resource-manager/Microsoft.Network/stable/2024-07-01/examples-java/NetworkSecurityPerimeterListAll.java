
/**
 * Samples for NetworkSecurityPerimeters List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * NetworkSecurityPerimeterListAll.json
     */
    /**
     * Sample code: NetworkSecurityPerimetersList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkSecurityPerimetersList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeters().list(null, null,
            com.azure.core.util.Context.NONE);
    }
}
