/** Samples for FileSystems ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().listByResourceGroup("rgQumulo", com.azure.core.util.Context.NONE);
    }
}
