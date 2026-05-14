
/**
 * Samples for FileShares ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShares_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesListByParentMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().listByResourceGroup("rgfileshares", com.azure.core.util.Context.NONE);
    }
}
