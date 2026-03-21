
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
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void scopeMapGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getScopeMaps().getWithResponse("myResourceGroup", "myRegistry", "myScopeMap",
            com.azure.core.util.Context.NONE);
    }
}
