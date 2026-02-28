
/**
 * Samples for ScopeMaps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ScopeMapList.json
     */
    /**
     * Sample code: ScopeMapList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getScopeMaps().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
