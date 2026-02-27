
/**
 * Samples for ScopeMaps Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ScopeMapGet.json
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
