
/**
 * Samples for TransitHub Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TransitHub_Get.json
     */
    /**
     * Sample code: TransitHub_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void transitHubGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.transitHubs().getWithResponse("rgopenapi", "TestMyCommunity", "TestThName",
            com.azure.core.util.Context.NONE);
    }
}
