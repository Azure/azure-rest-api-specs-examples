
/**
 * Samples for FileSystems List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * FileSystems_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void
        fileSystemsListBySubscriptionMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().list(com.azure.core.util.Context.NONE);
    }
}
