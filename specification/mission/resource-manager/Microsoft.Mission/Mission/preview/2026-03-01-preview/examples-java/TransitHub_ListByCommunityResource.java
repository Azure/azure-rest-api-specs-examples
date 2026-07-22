
/**
 * Samples for TransitHub ListByCommunityResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TransitHub_ListByCommunityResource.json
     */
    /**
     * Sample code: TransitHub_ListByCommunityResource.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        transitHubListByCommunityResource(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.transitHubs().listByCommunityResource("rgopenapi", "TestMyCommunity", com.azure.core.util.Context.NONE);
    }
}
