
/**
 * Samples for EnclaveConnection ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveConnection_ListByResourceGroup.json
     */
    /**
     * Sample code: EnclaveConnection_ListByResourceGroup.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveConnectionListByResourceGroup(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveConnections().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
