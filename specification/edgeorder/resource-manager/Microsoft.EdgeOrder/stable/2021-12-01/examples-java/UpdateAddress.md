```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.models.AddressResource;
import com.azure.resourcemanager.edgeorder.models.AddressType;
import com.azure.resourcemanager.edgeorder.models.ContactDetails;
import com.azure.resourcemanager.edgeorder.models.ShippingAddress;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourceProvider UpdateAddress. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateAddress.json
     */
    /**
     * Sample code: UpdateAddress.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void updateAddress(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        AddressResource resource =
            manager
                .resourceProviders()
                .getByResourceGroupWithResponse("TestRG", "TestAddressName2", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(
                mapOf(
                    "Hobby",
                    "Web Series Added",
                    "Name",
                    "Smile-Updated",
                    "WhatElse",
                    "Web Series Added",
                    "Work",
                    "Engineering"))
            .withShippingAddress(
                new ShippingAddress()
                    .withStreetAddress1("16 TOWNSEND STT")
                    .withStreetAddress2("UNIT 1")
                    .withCity("San Francisco")
                    .withStateOrProvince("CA")
                    .withCountry("US")
                    .withPostalCode("94107")
                    .withCompanyName("Microsoft")
                    .withAddressType(AddressType.NONE))
            .withContactDetails(
                new ContactDetails()
                    .withContactName("Petr Cech")
                    .withPhone("1234567890")
                    .withPhoneExtension("")
                    .withEmailList(Arrays.asList("ssemcr@microsoft.com")))
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.
