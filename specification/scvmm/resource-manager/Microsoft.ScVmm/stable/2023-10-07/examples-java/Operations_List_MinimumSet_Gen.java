
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Operations_List_MinimumSet_Gen.
     * json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void operationsListMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
