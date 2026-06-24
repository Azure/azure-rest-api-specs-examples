
/**
 * Samples for Usage List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/computeRPCommonExamples/Usage_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Usage_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void usageListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getUsages().list("4.0", com.azure.core.util.Context.NONE);
    }
}
