import com.azure.core.util.Context;

/** Samples for ScopeMaps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/ScopeMapList.json
     */
    /**
     * Sample code: ScopeMapList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getScopeMaps()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}
