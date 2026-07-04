
/**
 * Samples for NetworkProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkProfileDelete.json
     */
    /**
     * Sample code: Delete network profile.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkProfile(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkProfiles().delete("rg1", "networkProfile1", com.azure.core.util.Context.NONE);
    }
}
