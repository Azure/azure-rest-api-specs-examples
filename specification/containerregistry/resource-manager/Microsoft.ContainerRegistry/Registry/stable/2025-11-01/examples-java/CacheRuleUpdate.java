
import com.azure.resourcemanager.containerregistry.models.CacheRuleUpdateParameters;

/**
 * Samples for CacheRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CacheRuleUpdate.json
     */
    /**
     * Sample code: CacheRuleUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cacheRuleUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCacheRules().update("myResourceGroup", "myRegistry",
            "myCacheRule", new CacheRuleUpdateParameters().withCredentialSetResourceId("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
