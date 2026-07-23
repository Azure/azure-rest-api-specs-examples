
/**
 * Samples for Community GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Community_Get.json
     */
    /**
     * Sample code: Community_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communities().getByResourceGroupWithResponse("rgopenapi", "TestMyCommunity",
            com.azure.core.util.Context.NONE);
    }
}
