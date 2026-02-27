
/**
 * Samples for Tokens List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/TokenList.json
     */
    /**
     * Sample code: TokenList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTokens().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
