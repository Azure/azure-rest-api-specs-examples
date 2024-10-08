
import com.azure.resourcemanager.edgeorder.models.OrderItemResource;
import com.azure.resourcemanager.edgeorder.models.Preferences;
import com.azure.resourcemanager.edgeorder.models.TransportPreferences;
import com.azure.resourcemanager.edgeorder.models.TransportShipmentTypes;

/**
 * Samples for ResourceProvider UpdateOrderItem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
     */
    /**
     * Sample code: UpdateOrderItem.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void updateOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        OrderItemResource resource = manager.resourceProviders().getOrderItemByNameWithResponse("TestOrderItemName3",
            "YourResourceGroupName", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withPreferences(new Preferences().withTransportPreferences(
            new TransportPreferences().withPreferredShipmentType(TransportShipmentTypes.CUSTOMER_MANAGED))).apply();
    }
}
