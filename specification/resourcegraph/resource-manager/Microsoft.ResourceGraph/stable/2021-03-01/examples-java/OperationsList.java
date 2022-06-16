import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void operationsList(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager.operations().list(Context.NONE);
    }
}
