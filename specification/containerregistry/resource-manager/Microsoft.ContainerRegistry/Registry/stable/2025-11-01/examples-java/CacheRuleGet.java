
/**
 * Samples for CacheRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * CacheRuleGet.json
     */
    /**
     * Sample code: CacheRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cacheRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCacheRules().getWithResponse("myResourceGroup",
            "myRegistry", "myCacheRule", com.azure.core.util.Context.NONE);
    }
}
