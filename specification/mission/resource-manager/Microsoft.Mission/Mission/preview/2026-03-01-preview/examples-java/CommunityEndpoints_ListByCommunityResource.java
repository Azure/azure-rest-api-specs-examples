
/**
 * Samples for CommunityEndpoints ListByCommunityResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_ListByCommunityResource.json
     */
    /**
     * Sample code: CommunityEndpoints_ListByCommunityResource.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        communityEndpointsListByCommunityResource(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().listByCommunityResource("rgopenapi", "TestMyCommunity",
            com.azure.core.util.Context.NONE);
    }
}
