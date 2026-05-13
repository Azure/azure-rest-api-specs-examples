
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileShares_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesListBySubscriptionMinimumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().list(com.azure.core.util.Context.NONE);
    }
}
