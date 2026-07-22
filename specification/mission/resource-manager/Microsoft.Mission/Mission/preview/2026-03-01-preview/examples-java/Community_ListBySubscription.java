
/**
 * Samples for Community List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Community_ListBySubscription.json
     */
    /**
     * Sample code: Community_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communities().list(com.azure.core.util.Context.NONE);
    }
}
