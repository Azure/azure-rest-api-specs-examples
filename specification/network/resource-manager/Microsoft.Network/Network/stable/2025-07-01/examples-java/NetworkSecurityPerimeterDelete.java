
/**
 * Samples for NetworkSecurityPerimeters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityPerimeterDelete.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkSecurityPerimeterDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeters().delete("rg1", "testNSP1", null,
            com.azure.core.util.Context.NONE);
    }
}
