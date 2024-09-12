
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void operationsList(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
