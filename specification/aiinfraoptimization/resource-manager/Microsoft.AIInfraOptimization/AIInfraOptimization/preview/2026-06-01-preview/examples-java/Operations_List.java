
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/Operations_List.json
     */
    /**
     * Sample code: List the operations for the Microsoft.AIInfraOptimization resource provider.
     * 
     * @param manager Entry point to AIInfraOptimizationManager.
     */
    public static void listTheOperationsForTheMicrosoftAIInfraOptimizationResourceProvider(
        com.azure.resourcemanager.aiinfraoptimization.AIInfraOptimizationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
