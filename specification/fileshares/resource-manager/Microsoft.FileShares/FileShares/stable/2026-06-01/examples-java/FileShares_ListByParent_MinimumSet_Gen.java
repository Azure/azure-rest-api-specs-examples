
/**
 * Samples for FileShares ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_ListByParent_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileShares_ListByParent_MinimumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesListByParentMinimumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().listByResourceGroup("rgfileshares", com.azure.core.util.Context.NONE);
    }
}
