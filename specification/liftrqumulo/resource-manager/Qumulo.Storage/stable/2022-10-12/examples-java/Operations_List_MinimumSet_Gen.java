/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void operationsListMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
