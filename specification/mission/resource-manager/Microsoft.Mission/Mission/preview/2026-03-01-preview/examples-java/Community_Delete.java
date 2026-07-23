
/**
 * Samples for Community Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Community_Delete.json
     */
    /**
     * Sample code: Community_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communities().delete("rgopenapi", "TestMyCommunity", com.azure.core.util.Context.NONE);
    }
}
