/** Samples for Customers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerWithExpand.json
     */
    /**
     * Sample code: CustomerWithExpand.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void customerWithExpand(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .customers()
            .getWithResponse(
                "{billingAccountName}",
                "{customerName}",
                "enabledAzurePlans,resellers",
                com.azure.core.util.Context.NONE);
    }
}
