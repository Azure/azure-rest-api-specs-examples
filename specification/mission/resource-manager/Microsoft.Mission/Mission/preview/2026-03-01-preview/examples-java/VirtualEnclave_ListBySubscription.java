
/**
 * Samples for VirtualEnclave List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/VirtualEnclave_ListBySubscription.json
     */
    /**
     * Sample code: VirtualEnclave_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void virtualEnclaveListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.virtualEnclaves().list(com.azure.core.util.Context.NONE);
    }
}
