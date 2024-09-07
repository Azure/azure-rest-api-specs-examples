
import java.time.OffsetDateTime;

/**
 * Samples for BillingAccounts CancelPaymentTerms.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentTermsCancel.json
     */
    /**
     * Sample code: PaymentTermsCancel.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void paymentTermsCancel(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().cancelPaymentTerms(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            OffsetDateTime.parse("2023-01-05T22:39:34.2606750Z"), com.azure.core.util.Context.NONE);
    }
}
