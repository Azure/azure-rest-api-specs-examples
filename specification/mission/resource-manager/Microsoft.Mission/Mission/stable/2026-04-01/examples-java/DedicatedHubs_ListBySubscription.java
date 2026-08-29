
/**
 * Samples for DedicatedHub ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/DedicatedHubs_ListBySubscription.json
     */
    /**
     * Sample code: DedicatedHub_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void dedicatedHubListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.dedicatedHubs().listBySubscription("TestCommunity1", com.azure.core.util.Context.NONE);
    }
}
