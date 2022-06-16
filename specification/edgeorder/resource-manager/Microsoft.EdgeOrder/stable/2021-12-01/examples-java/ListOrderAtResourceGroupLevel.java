import com.azure.core.util.Context;

/** Samples for ResourceProvider ListOrderAtResourceGroupLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderAtResourceGroupLevel.json
     */
    /**
     * Sample code: ListOrderAtResourceGroupLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listOrderAtResourceGroupLevel(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listOrderAtResourceGroupLevel("TestRG", null, Context.NONE);
    }
}
