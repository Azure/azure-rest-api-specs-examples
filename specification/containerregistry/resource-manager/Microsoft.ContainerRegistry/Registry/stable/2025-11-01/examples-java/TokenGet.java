
/**
 * Samples for Tokens Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/TokenGet.json
     */
    /**
     * Sample code: TokenGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTokens().getWithResponse("myResourceGroup",
            "myRegistry", "myToken", com.azure.core.util.Context.NONE);
    }
}
