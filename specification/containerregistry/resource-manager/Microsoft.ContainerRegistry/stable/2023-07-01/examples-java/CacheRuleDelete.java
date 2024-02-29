
/**
 * Samples for CacheRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/
     * CacheRuleDelete.json
     */
    /**
     * Sample code: CacheRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cacheRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCacheRules().delete("myResourceGroup", "myRegistry",
            "myCacheRule", com.azure.core.util.Context.NONE);
    }
}
