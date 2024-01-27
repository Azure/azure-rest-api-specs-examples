
/** Samples for ScopeMaps Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/
     * stable/2023-07-01/examples/ScopeMapGet.json
     */
    /**
     * Sample code: ScopeMapGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getScopeMaps().getWithResponse("myResourceGroup",
            "myRegistry", "myScopeMap", com.azure.core.util.Context.NONE);
    }
}
