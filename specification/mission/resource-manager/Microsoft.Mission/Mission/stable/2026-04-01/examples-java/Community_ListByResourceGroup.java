
/**
 * Samples for Community ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Community_ListByResourceGroup.json
     */
    /**
     * Sample code: Community_ListByResourceGroup.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void communityListByResourceGroup(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communities().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
