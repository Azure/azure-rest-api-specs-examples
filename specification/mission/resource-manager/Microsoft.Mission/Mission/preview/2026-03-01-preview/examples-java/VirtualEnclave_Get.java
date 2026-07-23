
/**
 * Samples for VirtualEnclave GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/VirtualEnclave_Get.json
     */
    /**
     * Sample code: VirtualEnclave_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void virtualEnclaveGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.virtualEnclaves().getByResourceGroupWithResponse("rgopenapi", "TestMyEnclave",
            com.azure.core.util.Context.NONE);
    }
}
