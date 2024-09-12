
/**
 * Samples for FileSystems GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * FileSystems_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Get.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsGet(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().getByResourceGroupWithResponse("rgQumulo", "sihbehcisdqtqqyfiewiiaphgh",
            com.azure.core.util.Context.NONE);
    }
}
