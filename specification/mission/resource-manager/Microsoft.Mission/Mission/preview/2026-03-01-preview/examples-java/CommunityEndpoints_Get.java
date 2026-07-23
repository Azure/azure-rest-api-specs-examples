
/**
 * Samples for CommunityEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_Get.json
     */
    /**
     * Sample code: CommunityEndpoints_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityEndpointsGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().getWithResponse("rgopenapi", "TestMyCommunity", "TestMyCommunityEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
