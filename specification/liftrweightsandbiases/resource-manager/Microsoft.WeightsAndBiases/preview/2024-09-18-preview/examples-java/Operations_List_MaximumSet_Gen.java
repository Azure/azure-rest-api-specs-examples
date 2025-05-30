
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-18-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to WeightsAndBiasesManager.
     */
    public static void operationsListGeneratedByMaximumSetRule(
        com.azure.resourcemanager.weightsandbiases.WeightsAndBiasesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
