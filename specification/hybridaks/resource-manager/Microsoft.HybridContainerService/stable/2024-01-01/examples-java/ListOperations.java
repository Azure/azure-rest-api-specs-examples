
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        listOperations(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
