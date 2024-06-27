
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Operations_List_MaximumSet_Gen.
     * json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void operationsListMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
