
/**
 * Samples for FileSystems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-16/FileSystems_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsListBySubscriptionMaximumSet(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().list(com.azure.core.util.Context.NONE);
    }
}
