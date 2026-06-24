
/**
 * Samples for FileSystems GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-16/FileSystems_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Get_MaximumSet.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsGetMaximumSet(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().getByResourceGroupWithResponse("rgQumulo", "qumulo-fs-01",
            com.azure.core.util.Context.NONE);
    }
}
