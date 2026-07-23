
/**
 * Samples for TransitHub Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TransitHub_Delete.json
     */
    /**
     * Sample code: TransitHub_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void transitHubDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.transitHubs().delete("rgopenapi", "TestMyCommunity", "TestThName", com.azure.core.util.Context.NONE);
    }
}
