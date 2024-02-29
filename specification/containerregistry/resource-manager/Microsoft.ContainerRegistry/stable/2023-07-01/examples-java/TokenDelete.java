
/**
 * Samples for Tokens Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/
     * TokenDelete.json
     */
    /**
     * Sample code: TokenDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTokens().delete("myResourceGroup", "myRegistry",
            "myToken", com.azure.core.util.Context.NONE);
    }
}
