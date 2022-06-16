import com.azure.core.util.Context;

/** Samples for Tokens List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/TokenList.json
     */
    /**
     * Sample code: TokenList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getTokens()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}
