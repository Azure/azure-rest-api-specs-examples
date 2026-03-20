
/**
 * Samples for Tokens Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/TokenDelete.json
     */
    /**
     * Sample code: TokenDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void tokenDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getTokens().delete("myResourceGroup", "myRegistry", "myToken",
            com.azure.core.util.Context.NONE);
    }
}
