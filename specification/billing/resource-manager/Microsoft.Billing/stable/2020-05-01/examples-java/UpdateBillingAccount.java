import com.azure.resourcemanager.billing.models.AddressDetails;
import com.azure.resourcemanager.billing.models.BillingAccountUpdateRequest;

/** Samples for BillingAccounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingAccount.json
     */
    /**
     * Sample code: UpdateBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updateBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingAccounts()
            .update(
                "{billingAccountName}",
                new BillingAccountUpdateRequest()
                    .withDisplayName("Test Account")
                    .withSoldTo(
                        new AddressDetails()
                            .withFirstName("Test")
                            .withLastName("User")
                            .withCompanyName("Contoso")
                            .withAddressLine1("Test Address 1")
                            .withCity("Redmond")
                            .withRegion("WA")
                            .withCountry("US")
                            .withPostalCode("fakeTokenPlaceholder")),
                com.azure.core.util.Context.NONE);
    }
}
