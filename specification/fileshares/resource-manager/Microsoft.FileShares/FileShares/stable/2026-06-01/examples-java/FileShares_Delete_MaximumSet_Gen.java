
/**
 * Samples for FileShares Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShares_Delete_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileSharesDeleteMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().delete("rgfileshares", "fileshare", com.azure.core.util.Context.NONE);
    }
}
