
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/operations/stable/2024-04-01/examples/
     * ListOperations.json
     */
    /**
     * Sample code: Create cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
