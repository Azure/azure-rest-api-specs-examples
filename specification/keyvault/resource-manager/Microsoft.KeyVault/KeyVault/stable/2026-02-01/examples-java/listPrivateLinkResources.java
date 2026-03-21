
/**
 * Samples for PrivateLinkResources ListByVault.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listPrivateLinkResources.json
     */
    /**
     * Sample code: KeyVaultListPrivateLinkResources.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void keyVaultListPrivateLinkResources(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getPrivateLinkResources().listByVaultWithResponse("sample-group", "sample-vault",
            com.azure.core.util.Context.NONE);
    }
}
