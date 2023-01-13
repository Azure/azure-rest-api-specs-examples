import com.azure.resourcemanager.billing.fluent.models.BillingProfileInner;
import com.azure.resourcemanager.billing.models.AddressDetails;
import com.azure.resourcemanager.billing.models.AzurePlan;
import java.util.Arrays;

/** Samples for BillingProfiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
     */
    /**
     * Sample code: CreateBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void createBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingProfiles()
            .createOrUpdate(
                "{billingAccountName}",
                "{billingProfileName}",
                new BillingProfileInner()
                    .withDisplayName("Finance")
                    .withPoNumber("ABC12345")
                    .withBillTo(
                        new AddressDetails()
                            .withFirstName("Test")
                            .withLastName("User")
                            .withAddressLine1("Test Address 1")
                            .withCity("Redmond")
                            .withRegion("WA")
                            .withCountry("US")
                            .withPostalCode("fakeTokenPlaceholder"))
                    .withInvoiceEmailOptIn(true)
                    .withEnabledAzurePlans(
                        Arrays.asList(new AzurePlan().withSkuId("0001"), new AzurePlan().withSkuId("0002"))),
                com.azure.core.util.Context.NONE);
    }
}
