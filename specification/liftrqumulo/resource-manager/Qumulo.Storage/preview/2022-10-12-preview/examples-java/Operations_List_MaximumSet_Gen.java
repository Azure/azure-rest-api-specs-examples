/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void operationsListMaximumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
