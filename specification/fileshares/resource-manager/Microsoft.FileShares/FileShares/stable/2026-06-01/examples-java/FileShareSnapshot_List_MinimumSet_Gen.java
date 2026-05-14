
/**
 * Samples for FileShareSnapshots ListByFileShare.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShareSnapshot_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileShareSnapshot_List_MinimumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareSnapshotListMinimumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShareSnapshots().listByFileShare("rgfileshares", "testfileshare", com.azure.core.util.Context.NONE);
    }
}
