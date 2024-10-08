
import com.azure.resourcemanager.edgeorder.models.ReturnOrderItemDetails;

/**
 * Samples for ResourceProvider ReturnOrderItem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ReturnOrderItem.json
     */
    /**
     * Sample code: ReturnOrderItem.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void returnOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().returnOrderItem("TestOrderName4", "YourResourceGroupName",
            new ReturnOrderItemDetails().withReturnReason("Order returned"), com.azure.core.util.Context.NONE);
    }
}
