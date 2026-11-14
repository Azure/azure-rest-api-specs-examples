
/**
 * Samples for ScopeMaps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-15-preview/ScopeMapList.json
     */
    /**
     * Sample code: ScopeMapList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void scopeMapList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getScopeMaps().list("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
