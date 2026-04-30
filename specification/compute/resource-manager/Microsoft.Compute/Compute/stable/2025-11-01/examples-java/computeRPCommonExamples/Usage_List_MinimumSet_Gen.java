
/**
 * Samples for Usage List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/computeRPCommonExamples/Usage_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Usage_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void usageListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getUsages().list("_--", com.azure.core.util.Context.NONE);
    }
}
