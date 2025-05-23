
/**
 * Samples for Instances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-18-preview/Instances_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_ListBySubscription - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to WeightsAndBiasesManager.
     */
    public static void instancesListBySubscriptionGeneratedByMaximumSetRule(
        com.azure.resourcemanager.weightsandbiases.WeightsAndBiasesManager manager) {
        manager.instances().list(com.azure.core.util.Context.NONE);
    }
}
