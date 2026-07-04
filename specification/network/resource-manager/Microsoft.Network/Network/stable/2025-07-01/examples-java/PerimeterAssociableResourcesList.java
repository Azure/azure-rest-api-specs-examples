
/**
 * Samples for NetworkSecurityPerimeterAssociableResourceTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PerimeterAssociableResourcesList.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterAssociableResourceTypes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        networkSecurityPerimeterAssociableResourceTypes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAssociableResourceTypes().list("westus",
            com.azure.core.util.Context.NONE);
    }
}
