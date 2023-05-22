/** Samples for FileSystems GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsGetMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager
            .fileSystems()
            .getByResourceGroupWithResponse("rgQumulo", "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
