
/**
 * Samples for FileShareSnapshots ListByFileShare.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShareSnapshot_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShareSnapshot_List_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareSnapshotListMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShareSnapshots().listByFileShare("rgfileshares", "fileshare", com.azure.core.util.Context.NONE);
    }
}
