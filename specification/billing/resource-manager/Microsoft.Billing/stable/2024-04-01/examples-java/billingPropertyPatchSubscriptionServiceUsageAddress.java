
import com.azure.resourcemanager.billing.fluent.models.BillingPropertyInner;
import com.azure.resourcemanager.billing.models.BillingPropertyProperties;
import com.azure.resourcemanager.billing.models.BillingPropertyPropertiesSubscriptionServiceUsageAddress;

/**
 * Samples for BillingProperty Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPropertyPatchSubscriptionServiceUsageAddress.json
     */
    /**
     * Sample code: BillingPropertyPatchSubscriptionServiceUsageAddress.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingPropertyPatchSubscriptionServiceUsageAddress(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProperties()
            .updateWithResponse(
                new BillingPropertyInner().withProperties(new BillingPropertyProperties()
                    .withSubscriptionServiceUsageAddress(new BillingPropertyPropertiesSubscriptionServiceUsageAddress()
                        .withAddressLine1("Address line 1").withAddressLine2("Address line 2").withCity("City")
                        .withCountry("US").withFirstName("Jenny").withLastName("Doe").withMiddleName("Ann")
                        .withPostalCode("fakeTokenPlaceholder").withRegion("State"))),
                com.azure.core.util.Context.NONE);
    }
}
