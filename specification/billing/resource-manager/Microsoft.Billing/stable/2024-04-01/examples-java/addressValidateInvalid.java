
import com.azure.resourcemanager.billing.models.AddressDetails;

/**
 * Samples for Address Validate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/addressValidateInvalid.json
     */
    /**
     * Sample code: AddressValidateInvalid.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void addressValidateInvalid(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.address().validateWithResponse(new AddressDetails().withAddressLine1("1 Test").withCity("bellevue")
            .withCountry("us").withPostalCode("fakeTokenPlaceholder").withRegion("wa"),
            com.azure.core.util.Context.NONE);
    }
}
