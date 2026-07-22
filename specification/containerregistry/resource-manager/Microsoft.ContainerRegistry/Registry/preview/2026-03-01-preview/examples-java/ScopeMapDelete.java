
/**
 * Samples for ScopeMaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ScopeMapDelete.json
     */
    /**
     * Sample code: ScopeMapDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void scopeMapDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getScopeMaps().delete("myResourceGroup", "myRegistry", "myScopeMap",
            com.azure.core.util.Context.NONE);
    }
}
