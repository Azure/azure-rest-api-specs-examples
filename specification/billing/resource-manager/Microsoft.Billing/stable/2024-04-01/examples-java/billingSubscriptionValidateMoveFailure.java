
import com.azure.resourcemanager.billing.models.MoveBillingSubscriptionRequest;

/**
 * Samples for BillingSubscriptions ValidateMoveEligibility.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionValidateMoveFailure.json
     */
    /**
     * Sample code: BillingSubscriptionValidateMoveFailure.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionValidateMoveFailure(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().validateMoveEligibilityWithResponse(
            "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31",
            "6b96d3f2-9008-4a9d-912f-f87744185aa3",
            new MoveBillingSubscriptionRequest().withDestinationInvoiceSectionId(
                "/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
            com.azure.core.util.Context.NONE);
    }
}
