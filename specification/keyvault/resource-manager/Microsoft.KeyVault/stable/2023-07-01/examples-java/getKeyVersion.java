
/**
 * Samples for Keys GetVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/getKeyVersion.json
     */
    /**
     * Sample code: Get a key version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAKeyVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getKeys().getVersionWithResponse("sample-group", "sample-vault-name",
            "sample-key-name", "fd618d9519b74f9aae94ade66b876acc", com.azure.core.util.Context.NONE);
    }
}
