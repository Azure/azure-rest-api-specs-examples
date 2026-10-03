
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void operationsListMaximumSetGen(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
