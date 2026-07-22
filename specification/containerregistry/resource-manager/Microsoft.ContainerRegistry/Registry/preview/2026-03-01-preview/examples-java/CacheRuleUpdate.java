
import com.azure.resourcemanager.containerregistry.models.CacheRuleUpdateParameters;

/**
 * Samples for CacheRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CacheRuleUpdate.json
     */
    /**
     * Sample code: CacheRuleUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleUpdate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().update("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleUpdateParameters().withCredentialSetResourceId("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
