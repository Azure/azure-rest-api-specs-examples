import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.fluent.models.AddressProperties;
import com.azure.resourcemanager.edgeorder.models.ContactDetails;
import com.azure.resourcemanager.edgeorder.models.OrderItemResource;
import com.azure.resourcemanager.edgeorder.models.Preferences;
import com.azure.resourcemanager.edgeorder.models.TransportPreferences;
import com.azure.resourcemanager.edgeorder.models.TransportShipmentTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourceProvider UpdateOrderItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
     */
    /**
     * Sample code: UpdateOrderItem.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void updateOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        OrderItemResource resource =
            manager
                .resourceProviders()
                .getOrderItemByNameWithResponse("TestOrderItemName01", "TestRG", null, Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("ant", "insect", "pigeon", "bird", "tiger", "animal"))
            .withForwardAddress(
                new AddressProperties()
                    .withContactDetails(
                        new ContactDetails()
                            .withContactName("Updated contact name")
                            .withPhone("2222200000")
                            .withEmailList(Arrays.asList("testemail@microsoft.com"))))
            .withPreferences(
                new Preferences()
                    .withTransportPreferences(
                        new TransportPreferences().withPreferredShipmentType(TransportShipmentTypes.CUSTOMER_MANAGED)))
            .apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
