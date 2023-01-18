/** Samples for PrivateLinkScopes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-03-25-preview/examples/PrivateLinkScopesDelete.json
     */
    /**
     * Sample code: PrivateLinkScopesDelete.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopesDelete(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .privateLinkScopes()
            .delete("my-resource-group", "my-privatelinkscope", com.azure.core.util.Context.NONE);
    }
}
