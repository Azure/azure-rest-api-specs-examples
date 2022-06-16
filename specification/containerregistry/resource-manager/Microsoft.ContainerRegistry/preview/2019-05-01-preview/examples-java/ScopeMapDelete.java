import com.azure.core.util.Context;

/** Samples for ScopeMaps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/ScopeMapDelete.json
     */
    /**
     * Sample code: ScopeMapDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getScopeMaps()
            .delete("myResourceGroup", "myRegistry", "myScopeMap", Context.NONE);
    }
}
