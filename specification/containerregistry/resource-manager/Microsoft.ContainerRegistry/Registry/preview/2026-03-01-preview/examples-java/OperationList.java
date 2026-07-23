
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/OperationList.json
     */
    /**
     * Sample code: OperationList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void operationList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
