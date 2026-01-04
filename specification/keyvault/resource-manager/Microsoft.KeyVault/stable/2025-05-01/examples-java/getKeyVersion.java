
/**
 * Samples for Keys GetVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/getKeyVersion.json
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
