
/**
 * Samples for NetworkProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkProfileListAll.json
     */
    /**
     * Sample code: List all network profiles.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNetworkProfiles(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkProfiles().list(com.azure.core.util.Context.NONE);
    }
}
