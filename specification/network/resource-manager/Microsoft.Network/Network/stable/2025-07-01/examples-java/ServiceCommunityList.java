
/**
 * Samples for BgpServiceCommunities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceCommunityList.json
     */
    /**
     * Sample code: ServiceCommunityList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void serviceCommunityList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBgpServiceCommunities().list(com.azure.core.util.Context.NONE);
    }
}
