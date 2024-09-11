
/**
 * Samples for FileSystems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * FileSystems_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsDeleteMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().delete("rgQumulo", "jgtskkiplquyrlkaxvhdg", com.azure.core.util.Context.NONE);
    }
}
