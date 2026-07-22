
/**
 * Samples for EnclaveEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveEndpoints_Get.json
     */
    /**
     * Sample code: EnclaveEndpoints_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void enclaveEndpointsGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().getWithResponse("rgopenapi", "TestMyEnclave", "TestMyEnclaveEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
