
/**
 * Samples for DedicatedHub Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DedicatedHubs_Delete.json
     */
    /**
     * Sample code: DedicatedHub_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void dedicatedHubDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.dedicatedHubs().delete("TestResourceGroup", "TestCommunity", "TestDedicatedHub",
            com.azure.core.util.Context.NONE);
    }
}
