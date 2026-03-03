
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void operationsListMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
