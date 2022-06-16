import com.azure.core.util.Context;

/** Samples for ResourceProvider ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listOperations(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listOperations(Context.NONE);
    }
}
