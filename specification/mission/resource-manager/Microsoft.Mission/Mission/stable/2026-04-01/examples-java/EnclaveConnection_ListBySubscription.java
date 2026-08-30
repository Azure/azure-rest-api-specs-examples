
/**
 * Samples for EnclaveConnection List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/EnclaveConnection_ListBySubscription.json
     */
    /**
     * Sample code: EnclaveConnection_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveConnectionListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveConnections().list(com.azure.core.util.Context.NONE);
    }
}
