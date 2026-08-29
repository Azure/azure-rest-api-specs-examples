
/**
 * Samples for EnclaveEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/EnclaveEndpoints_Delete.json
     */
    /**
     * Sample code: EnclaveEndpoints_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void enclaveEndpointsDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().delete("rgopenapi", "TestMyEnclave", "TestMyEnclaveEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
