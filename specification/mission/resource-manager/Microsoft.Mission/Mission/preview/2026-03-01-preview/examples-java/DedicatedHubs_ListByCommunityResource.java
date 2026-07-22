
/**
 * Samples for DedicatedHub ListByCommunityResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DedicatedHubs_ListByCommunityResource.json
     */
    /**
     * Sample code: DedicatedHub_ListByCommunityResource.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        dedicatedHubListByCommunityResource(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.dedicatedHubs().listByCommunityResource("TestResourceGroup", "TestCommunity",
            com.azure.core.util.Context.NONE);
    }
}
