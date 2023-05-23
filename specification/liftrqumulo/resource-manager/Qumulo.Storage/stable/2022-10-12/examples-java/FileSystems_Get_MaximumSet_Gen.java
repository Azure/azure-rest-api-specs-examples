/** Samples for FileSystems GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsGetMaximumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager
            .fileSystems()
            .getByResourceGroupWithResponse(
                "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", com.azure.core.util.Context.NONE);
    }
}
