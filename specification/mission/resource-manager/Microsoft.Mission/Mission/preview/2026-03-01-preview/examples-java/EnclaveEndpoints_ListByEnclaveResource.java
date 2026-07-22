
/**
 * Samples for EnclaveEndpoints ListByEnclaveResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveEndpoints_ListByEnclaveResource.json
     */
    /**
     * Sample code: EnclaveEndpoints_ListByEnclaveResource.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveEndpointsListByEnclaveResource(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().listByEnclaveResource("rgopenapi", "TestMyEnclave",
            com.azure.core.util.Context.NONE);
    }
}
