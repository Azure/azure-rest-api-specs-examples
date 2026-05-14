
/**
 * Samples for FileShareSnapshots GetFileShareSnapshot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShareSnapshot_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShareSnapshot_Get_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareSnapshotGetMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShareSnapshots().getFileShareSnapshotWithResponse("rgfileshares", "fileshare",
            "testfilesharesnapshot", com.azure.core.util.Context.NONE);
    }
}
