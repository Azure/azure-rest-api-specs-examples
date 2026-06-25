
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-16/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void operationsListMaximumSet(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
