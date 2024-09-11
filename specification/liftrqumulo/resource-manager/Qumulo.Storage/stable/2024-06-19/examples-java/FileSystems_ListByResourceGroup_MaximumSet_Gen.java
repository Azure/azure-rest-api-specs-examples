
/**
 * Samples for FileSystems ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * FileSystems_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListByResourceGroup.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsListByResourceGroup(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().listByResourceGroup("rgQumulo", com.azure.core.util.Context.NONE);
    }
}
