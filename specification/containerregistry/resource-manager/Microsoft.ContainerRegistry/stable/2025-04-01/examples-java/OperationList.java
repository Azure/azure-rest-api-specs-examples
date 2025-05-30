
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * OperationList.json
     */
    /**
     * Sample code: OperationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void operationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
