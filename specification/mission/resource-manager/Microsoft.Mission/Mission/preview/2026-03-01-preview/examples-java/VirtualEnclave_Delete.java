
/**
 * Samples for VirtualEnclave Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/VirtualEnclave_Delete.json
     */
    /**
     * Sample code: VirtualEnclave_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void virtualEnclaveDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.virtualEnclaves().delete("rgopenapi", "TestMyEnclave", com.azure.core.util.Context.NONE);
    }
}
