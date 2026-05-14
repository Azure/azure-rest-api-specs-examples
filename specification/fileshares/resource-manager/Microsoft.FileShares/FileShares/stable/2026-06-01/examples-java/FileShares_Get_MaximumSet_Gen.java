
/**
 * Samples for FileShares GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShares_Get_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileSharesGetMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().getByResourceGroupWithResponse("rgfileshares", "fileshare",
            com.azure.core.util.Context.NONE);
    }
}
