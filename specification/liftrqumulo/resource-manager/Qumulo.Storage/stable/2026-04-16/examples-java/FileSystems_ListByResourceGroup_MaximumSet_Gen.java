
/**
 * Samples for FileSystems ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-16/FileSystems_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void
        fileSystemsListByResourceGroupMaximumSet(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().listByResourceGroup("rgQumulo", com.azure.core.util.Context.NONE);
    }
}
