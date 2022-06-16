import com.azure.core.util.Context;

/** Samples for Instructions ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InstructionsListByBillingProfile.json
     */
    /**
     * Sample code: InstructionsListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void instructionsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.instructions().listByBillingProfile("{billingAccountName}", "{billingProfileName}", Context.NONE);
    }
}
