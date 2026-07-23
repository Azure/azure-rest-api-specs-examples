
/**
 * Samples for Tokens Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TokenGet.json
     */
    /**
     * Sample code: TokenGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void tokenGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getTokens().getWithResponse("myResourceGroup", "myRegistry", "myToken",
            com.azure.core.util.Context.NONE);
    }
}
