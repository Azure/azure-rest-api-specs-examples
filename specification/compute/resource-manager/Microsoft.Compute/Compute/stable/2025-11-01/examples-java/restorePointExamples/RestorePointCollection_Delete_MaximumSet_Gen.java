
/**
 * Samples for RestorePointCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePointCollection_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePointCollection_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        restorePointCollectionDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().delete("rgcompute", "aaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
