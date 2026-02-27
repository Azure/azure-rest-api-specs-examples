
/**
 * Samples for ScopeMaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ScopeMapDelete.json
     */
    /**
     * Sample code: ScopeMapDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getScopeMaps().delete("myResourceGroup", "myRegistry",
            "myScopeMap", com.azure.core.util.Context.NONE);
    }
}
