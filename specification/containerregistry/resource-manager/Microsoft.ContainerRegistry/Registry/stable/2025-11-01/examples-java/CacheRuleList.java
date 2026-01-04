
/**
 * Samples for CacheRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * CacheRuleList.json
     */
    /**
     * Sample code: CacheRuleList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cacheRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCacheRules().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
