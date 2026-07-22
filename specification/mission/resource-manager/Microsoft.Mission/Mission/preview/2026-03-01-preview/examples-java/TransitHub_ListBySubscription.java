
/**
 * Samples for TransitHub ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TransitHub_ListBySubscription.json
     */
    /**
     * Sample code: TransitHub_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void transitHubListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.transitHubs().listBySubscription("TestMyCommunity", com.azure.core.util.Context.NONE);
    }
}
