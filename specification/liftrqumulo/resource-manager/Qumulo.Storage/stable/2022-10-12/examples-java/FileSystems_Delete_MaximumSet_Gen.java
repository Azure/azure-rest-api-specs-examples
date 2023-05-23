/** Samples for FileSystems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsDeleteMaximumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().delete("rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", com.azure.core.util.Context.NONE);
    }
}
