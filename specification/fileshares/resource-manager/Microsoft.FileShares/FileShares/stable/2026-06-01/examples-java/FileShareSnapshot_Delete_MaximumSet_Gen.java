
/**
 * Samples for FileShareSnapshots DeleteFileShareSnapshot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShareSnapshot_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShareSnapshot_Delete_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileShareSnapshotDeleteMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShareSnapshots().deleteFileShareSnapshot("rgfileshares", "fileshare", "testfilesharesnapshot",
            com.azure.core.util.Context.NONE);
    }
}
