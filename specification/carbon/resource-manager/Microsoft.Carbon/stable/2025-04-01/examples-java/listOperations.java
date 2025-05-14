
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/listOperations.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to CarbonOptimizationManager.
     */
    public static void operationsList(com.azure.resourcemanager.carbonoptimization.CarbonOptimizationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
