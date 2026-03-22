
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listOperations.json
     */
    /**
     * Sample code: Lists available Rest API operations.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listsAvailableRestAPIOperations(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
