import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listOperations(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.operations().list(Context.NONE);
    }
}
