
import com.azure.resourcemanager.billing.models.BillingAccountPatch;
import com.azure.resourcemanager.billing.models.BillingAccountProperties;
import com.azure.resourcemanager.billing.models.BillingAccountPropertiesSoldTo;

/**
 * Samples for BillingAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsUpdate.json
     */
    /**
     * Sample code: BillingAccountsUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().update(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            new BillingAccountPatch().withProperties(new BillingAccountProperties().withDisplayName("Test Account")
                .withSoldTo(new BillingAccountPropertiesSoldTo().withAddressLine1("1 Microsoft Way").withCity("Redmond")
                    .withCompanyName("Contoso").withCountry("US").withPostalCode("fakeTokenPlaceholder")
                    .withRegion("WA"))),
            com.azure.core.util.Context.NONE);
    }
}
