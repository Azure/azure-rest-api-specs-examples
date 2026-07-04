
/**
 * Samples for AvailableServiceAliases List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailableServiceAliasesList.json
     */
    /**
     * Sample code: Get available service aliases.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableServiceAliases(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailableServiceAliases().list("westcentralus", com.azure.core.util.Context.NONE);
    }
}
