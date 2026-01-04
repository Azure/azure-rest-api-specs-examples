
/**
 * Samples for Vaults GetDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/getDeletedVault.json
     */
    /**
     * Sample code: Retrieve a deleted vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveADeletedVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().getDeletedWithResponse("sample-vault", "westus",
            com.azure.core.util.Context.NONE);
    }
}
