/** Samples for FileSystems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListBySubscription_MaximumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().list(com.azure.core.util.Context.NONE);
    }
}
