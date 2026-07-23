
/**
 * Samples for CommunityEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_Delete.json
     */
    /**
     * Sample code: CommunityEndpoints_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityEndpointsDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().delete("rgopenapi", "TestMyCommunity", "TestMyCommunityEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
