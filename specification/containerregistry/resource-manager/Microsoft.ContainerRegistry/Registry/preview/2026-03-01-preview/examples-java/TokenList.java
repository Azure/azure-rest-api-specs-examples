
/**
 * Samples for Tokens List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/TokenList.json
     */
    /**
     * Sample code: TokenList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void tokenList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getTokens().list("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
