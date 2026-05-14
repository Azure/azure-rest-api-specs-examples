
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShares_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesListBySubscriptionMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().list(com.azure.core.util.Context.NONE);
    }
}
