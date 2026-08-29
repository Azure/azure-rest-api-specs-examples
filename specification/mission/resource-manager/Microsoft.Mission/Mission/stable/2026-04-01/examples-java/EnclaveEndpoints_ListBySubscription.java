
/**
 * Samples for EnclaveEndpoints ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/EnclaveEndpoints_ListBySubscription.json
     */
    /**
     * Sample code: EnclaveEndpoints_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveEndpointsListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().listBySubscription("TestMyEnclave", com.azure.core.util.Context.NONE);
    }
}
