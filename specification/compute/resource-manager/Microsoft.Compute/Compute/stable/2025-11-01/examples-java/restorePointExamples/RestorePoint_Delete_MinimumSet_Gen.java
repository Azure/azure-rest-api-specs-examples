
/**
 * Samples for RestorePoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePoint_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: RestorePoint_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void restorePointDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().delete("rgcompute", "aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
