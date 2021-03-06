import com.azure.resourcemanager.edgeorder.fluent.models.AddressProperties;
import com.azure.resourcemanager.edgeorder.models.AddressDetails;
import com.azure.resourcemanager.edgeorder.models.AddressType;
import com.azure.resourcemanager.edgeorder.models.ContactDetails;
import com.azure.resourcemanager.edgeorder.models.HierarchyInformation;
import com.azure.resourcemanager.edgeorder.models.OrderItemDetails;
import com.azure.resourcemanager.edgeorder.models.OrderItemType;
import com.azure.resourcemanager.edgeorder.models.Preferences;
import com.azure.resourcemanager.edgeorder.models.ProductDetails;
import com.azure.resourcemanager.edgeorder.models.ShippingAddress;
import com.azure.resourcemanager.edgeorder.models.TransportPreferences;
import com.azure.resourcemanager.edgeorder.models.TransportShipmentTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourceProvider CreateOrderItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateOrderItem.json
     */
    /**
     * Sample code: CreateOrderItem.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void createOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .defineOrderItem("TestOrderItemName01")
            .withRegion("westus")
            .withExistingResourceGroup("TestRG")
            .withOrderItemDetails(
                new OrderItemDetails()
                    .withProductDetails(
                        new ProductDetails()
                            .withHierarchyInformation(
                                new HierarchyInformation()
                                    .withProductFamilyName("AzureStackEdge")
                                    .withProductLineName("AzureStackEdge")
                                    .withProductName("AzureStackEdgeGPU")
                                    .withConfigurationName("AzureStackEdgeGPU")))
                    .withOrderItemType(OrderItemType.PURCHASE)
                    .withPreferences(
                        new Preferences()
                            .withTransportPreferences(
                                new TransportPreferences()
                                    .withPreferredShipmentType(TransportShipmentTypes.MICROSOFT_MANAGED))))
            .withAddressDetails(
                new AddressDetails()
                    .withForwardAddress(
                        new AddressProperties()
                            .withShippingAddress(
                                new ShippingAddress()
                                    .withStreetAddress1("16 TOWNSEND ST")
                                    .withStreetAddress2("UNIT 1")
                                    .withCity("San Francisco")
                                    .withStateOrProvince("CA")
                                    .withCountry("US")
                                    .withPostalCode("94107")
                                    .withZipExtendedCode("1")
                                    .withCompanyName("Microsoft")
                                    .withAddressType(AddressType.RESIDENTIAL))
                            .withContactDetails(
                                new ContactDetails()
                                    .withContactName("164 TOWNSEND ST")
                                    .withPhone("3213131190")
                                    .withEmailList(
                                        Arrays.asList("ssemmail@microsoft.com", "vishwamdir@microsoft.com")))))
            .withOrderId(
                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/TestRG/providers/Microsoft.EdgeOrder/locations/westus/orders/TestOrderItemName01")
            .withTags(mapOf("carrot", "vegetable", "mango", "fruit"))
            .create();
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
