
/**
 * Samples for NetworkInterfaces ListCloudServiceNetworkInterfaces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceNetworkInterfaceList.json
     */
    /**
     * Sample code: List cloud service network interfaces.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listCloudServiceNetworkInterfaces(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().listCloudServiceNetworkInterfaces("rg1", "cs1",
            com.azure.core.util.Context.NONE);
    }
}
