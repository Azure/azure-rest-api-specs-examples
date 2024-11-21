
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/ListOperations.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        operationsList(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
