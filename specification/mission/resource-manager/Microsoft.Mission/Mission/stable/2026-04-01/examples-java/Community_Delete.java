
/**
 * Samples for Community Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Community_Delete.json
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
