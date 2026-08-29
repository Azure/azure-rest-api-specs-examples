
/**
 * Samples for TransitHub Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/TransitHub_Delete.json
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
