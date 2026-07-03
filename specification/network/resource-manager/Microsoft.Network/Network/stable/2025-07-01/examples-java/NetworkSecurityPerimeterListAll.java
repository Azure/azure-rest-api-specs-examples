
/**
 * Samples for NetworkSecurityPerimeters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityPerimeterListAll.json
     */
    /**
     * Sample code: NetworkSecurityPerimetersList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkSecurityPerimetersList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeters().list(null, null, com.azure.core.util.Context.NONE);
    }
}
