
/**
 * Samples for CommunityEndpoints ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_ListBySubscription.json
     */
    /**
     * Sample code: CommunityEndpoints_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        communityEndpointsListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().listBySubscription("TestMyCommunity", com.azure.core.util.Context.NONE);
    }
}
