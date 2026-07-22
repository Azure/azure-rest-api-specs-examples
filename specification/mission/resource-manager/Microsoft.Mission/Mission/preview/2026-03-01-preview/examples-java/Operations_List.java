
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void operationsList(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
