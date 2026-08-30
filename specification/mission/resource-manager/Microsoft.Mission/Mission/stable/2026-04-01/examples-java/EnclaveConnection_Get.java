
/**
 * Samples for EnclaveConnection GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/EnclaveConnection_Get.json
     */
    /**
     * Sample code: EnclaveConnection_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void enclaveConnectionGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveConnections().getByResourceGroupWithResponse("rgopenapi", "TestMyEnclaveConnection",
            com.azure.core.util.Context.NONE);
    }
}
