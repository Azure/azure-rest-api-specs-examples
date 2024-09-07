
import com.azure.resourcemanager.billing.models.BillingAccountPatch;
import com.azure.resourcemanager.billing.models.BillingAccountProperties;
import com.azure.resourcemanager.billing.models.BillingAccountPropertiesEnrollmentDetails;

/**
 * Samples for BillingAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingAccountUpdateWithPONumber.json
     */
    /**
     * Sample code: BillingAccountUpdateWithPONumber.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountUpdateWithPONumber(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().update("6575495",
            new BillingAccountPatch().withProperties(new BillingAccountProperties()
                .withEnrollmentDetails(new BillingAccountPropertiesEnrollmentDetails().withPoNumber("poNumber123"))),
            com.azure.core.util.Context.NONE);
    }
}
