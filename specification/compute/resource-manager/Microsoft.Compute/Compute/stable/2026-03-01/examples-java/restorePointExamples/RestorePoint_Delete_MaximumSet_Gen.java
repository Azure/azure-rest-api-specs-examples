
/**
 * Samples for RestorePoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePoint_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePoint_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void restorePointDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", "a",
            com.azure.core.util.Context.NONE);
    }
}
