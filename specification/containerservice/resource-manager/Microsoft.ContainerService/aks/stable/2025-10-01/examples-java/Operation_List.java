
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * Operation_List.json
     */
    /**
     * Sample code: List available operations for the container service resource provider.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAvailableOperationsForTheContainerServiceResourceProvider(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
