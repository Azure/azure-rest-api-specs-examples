import com.azure.resourcemanager.databoxedge.fluent.models.OrderInner;
import com.azure.resourcemanager.databoxedge.models.Address;
import com.azure.resourcemanager.databoxedge.models.ContactDetails;
import java.util.Arrays;

/** Samples for Orders CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderPut.json
     */
    /**
     * Sample code: OrderPut.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void orderPut(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .orders()
            .createOrUpdate(
                "testedgedevice",
                "GroupForEdgeAutomation",
                new OrderInner()
                    .withContactInformation(
                        new ContactDetails()
                            .withContactPerson("John Mcclane")
                            .withCompanyName("Microsoft")
                            .withPhone("(800) 426-9400")
                            .withEmailList(Arrays.asList("john@microsoft.com")))
                    .withShippingAddress(
                        new Address()
                            .withAddressLine1("Microsoft Corporation")
                            .withAddressLine2("One Microsoft Way")
                            .withAddressLine3("Redmond")
                            .withPostalCode("fakeTokenPlaceholder")
                            .withCity("WA")
                            .withState("WA")
                            .withCountry("USA")),
                com.azure.core.util.Context.NONE);
    }
}
