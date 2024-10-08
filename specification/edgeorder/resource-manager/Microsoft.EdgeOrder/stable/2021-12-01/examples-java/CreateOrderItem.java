
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

/**
 * Samples for ResourceProvider CreateOrderItem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateOrderItem.json
     */
    /**
     * Sample code: CreateOrderItem.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void createOrderItem(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().defineOrderItem("TestOrderItemName2").withRegion("eastus")
            .withExistingResourceGroup("YourResourceGroupName")
            .withOrderItemDetails(new OrderItemDetails()
                .withProductDetails(new ProductDetails().withHierarchyInformation(new HierarchyInformation()
                    .withProductFamilyName("azurestackedge").withProductLineName("azurestackedge")
                    .withProductName("azurestackedgegpu").withConfigurationName("edgep_base")))
                .withOrderItemType(OrderItemType.PURCHASE)
                .withPreferences(new Preferences().withTransportPreferences(
                    new TransportPreferences().withPreferredShipmentType(TransportShipmentTypes.MICROSOFT_MANAGED))))
            .withAddressDetails(new AddressDetails().withForwardAddress(new AddressProperties()
                .withShippingAddress(new ShippingAddress().withStreetAddress1("16 TOWNSEND ST")
                    .withStreetAddress2("UNIT 1").withCity("San Francisco").withStateOrProvince("CA").withCountry("US")
                    .withPostalCode("fakeTokenPlaceholder").withCompanyName("Microsoft")
                    .withAddressType(AddressType.NONE))
                .withContactDetails(new ContactDetails().withContactName("XXXX XXXX").withPhone("0000000000")
                    .withPhoneExtension("").withEmailList(Arrays.asList("xxxx@xxxx.xxx")))))
            .withOrderId(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName2")
            .create();
    }
}
