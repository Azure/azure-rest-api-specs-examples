/** Samples for BillingProperty Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProperty.json
     */
    /**
     * Sample code: BillingProperty.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProperty(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProperties().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
