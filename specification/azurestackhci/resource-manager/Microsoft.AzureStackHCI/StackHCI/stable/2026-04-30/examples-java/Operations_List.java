
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listTheOperationsForTheProvider(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
