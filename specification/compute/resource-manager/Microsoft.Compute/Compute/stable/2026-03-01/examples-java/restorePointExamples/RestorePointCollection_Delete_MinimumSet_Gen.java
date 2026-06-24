
/**
 * Samples for RestorePointCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePointCollection_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: RestorePointCollection_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        restorePointCollectionDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().delete("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
