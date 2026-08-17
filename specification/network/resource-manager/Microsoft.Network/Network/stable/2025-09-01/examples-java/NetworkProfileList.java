
/**
 * Samples for NetworkProfiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkProfileList.json
     */
    /**
     * Sample code: List resource group network profiles.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listResourceGroupNetworkProfiles(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkProfiles().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
