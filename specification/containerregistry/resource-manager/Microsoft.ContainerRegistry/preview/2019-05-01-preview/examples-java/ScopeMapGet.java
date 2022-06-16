import com.azure.core.util.Context;

/** Samples for ScopeMaps Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/ScopeMapGet.json
     */
    /**
     * Sample code: ScopeMapGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getScopeMaps()
            .getWithResponse("myResourceGroup", "myRegistry", "myScopeMap", Context.NONE);
    }
}
