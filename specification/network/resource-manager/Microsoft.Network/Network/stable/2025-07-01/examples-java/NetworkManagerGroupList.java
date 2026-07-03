
/**
 * Samples for NetworkGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerGroupList.json
     */
    /**
     * Sample code: NetworkGroupsList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkGroupsList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkGroups().list("rg1", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
