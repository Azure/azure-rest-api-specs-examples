import com.azure.resourcemanager.billing.models.AddressDetails;

/** Samples for Address Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AddressValid.json
     */
    /**
     * Sample code: AddressValid.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void addressValid(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .address()
            .validateWithResponse(
                new AddressDetails()
                    .withAddressLine1("1 Test Address")
                    .withCity("bellevue")
                    .withRegion("wa")
                    .withCountry("us")
                    .withPostalCode("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
