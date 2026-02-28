
/**
 * Samples for CacheRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CacheRuleDelete.json
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
