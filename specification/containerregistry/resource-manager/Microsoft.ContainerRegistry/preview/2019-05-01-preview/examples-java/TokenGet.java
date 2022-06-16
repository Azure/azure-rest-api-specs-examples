import com.azure.core.util.Context;

/** Samples for Tokens Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/TokenGet.json
     */
    /**
     * Sample code: TokenGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getTokens()
            .getWithResponse("myResourceGroup", "myRegistry", "myToken", Context.NONE);
    }
}
