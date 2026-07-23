
/**
 * Samples for DedicatedHub Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DedicatedHubs_Get.json
     */
    /**
     * Sample code: DedicatedHub_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void dedicatedHubGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.dedicatedHubs().getWithResponse("TestResourceGroup", "TestCommunity", "TestDedicatedHub",
            com.azure.core.util.Context.NONE);
    }
}
