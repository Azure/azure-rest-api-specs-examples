
/**
 * Samples for ConnectedRegistries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ConnectedRegistryList.json
     */
    /**
     * Sample code: ConnectedRegistryList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectedRegistryList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getConnectedRegistries().list("myResourceGroup",
            "myRegistry", null, com.azure.core.util.Context.NONE);
    }
}
