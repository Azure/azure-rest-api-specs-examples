/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ListSkus.json
     */
    /**
     * Sample code: ListSkus.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void listSkus(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.skus().list(null, com.azure.core.util.Context.NONE);
    }
}
