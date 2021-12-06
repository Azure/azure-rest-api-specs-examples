Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.models.AddressDetails;

/** Samples for Address Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AddressInvalid.json
     */
    /**
     * Sample code: AddressInvalid.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void addressInvalid(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .address()
            .validateWithResponse(
                new AddressDetails()
                    .withAddressLine1("1 Test")
                    .withCity("bellevue")
                    .withRegion("wa")
                    .withCountry("us")
                    .withPostalCode("12345"),
                Context.NONE);
    }
}
```
