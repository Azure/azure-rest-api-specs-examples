
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        operationsListMinimumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
