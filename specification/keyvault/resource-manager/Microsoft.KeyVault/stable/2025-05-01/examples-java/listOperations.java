
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/listOperations.json
     */
    /**
     * Sample code: Lists available Rest API operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAvailableRestAPIOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
