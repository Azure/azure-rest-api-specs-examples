
/**
 * Samples for Vaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/deleteVault.json
     */
    /**
     * Sample code: Delete a vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().deleteWithResponse("sample-resource-group", "sample-vault",
            com.azure.core.util.Context.NONE);
    }
}
