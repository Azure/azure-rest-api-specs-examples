Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.edgeorder.models.AddressType;
import com.azure.resourcemanager.edgeorder.models.ContactDetails;
import com.azure.resourcemanager.edgeorder.models.ShippingAddress;
import java.util.Arrays;

/** Samples for ResourceProvider CreateAddress. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateAddress.json
     */
    /**
     * Sample code: CreateAddress.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void createAddress(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .defineAddress("TestMSAddressName")
            .withRegion("westus")
            .withExistingResourceGroup("TestRG")
            .withContactDetails(
                new ContactDetails()
                    .withContactName("Petr Cech")
                    .withPhone("1234567890")
                    .withPhoneExtension("")
                    .withEmailList(Arrays.asList("testemail@microsoft.com")))
            .withShippingAddress(
                new ShippingAddress()
                    .withStreetAddress1("16 TOWNSEND ST")
                    .withStreetAddress2("UNIT 1")
                    .withCity("San Francisco")
                    .withStateOrProvince("CA")
                    .withCountry("US")
                    .withPostalCode("94107")
                    .withCompanyName("Microsoft")
                    .withAddressType(AddressType.NONE))
            .create();
    }
}
```
