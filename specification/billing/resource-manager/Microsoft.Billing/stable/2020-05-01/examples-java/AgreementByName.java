/** Samples for Agreements Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AgreementByName.json
     */
    /**
     * Sample code: AgreementByName.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void agreementByName(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .agreements()
            .getWithResponse("{billingAccountName}", "{agreementName}", null, com.azure.core.util.Context.NONE);
    }
}
