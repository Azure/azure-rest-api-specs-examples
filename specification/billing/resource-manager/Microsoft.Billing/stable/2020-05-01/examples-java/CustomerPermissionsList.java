/** Samples for BillingPermissions ListByCustomer. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPermissionsList.json
     */
    /**
     * Sample code: BillingProfilePermissionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfilePermissionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingPermissions()
            .listByCustomer("{billingAccountName}", "{customerName}", com.azure.core.util.Context.NONE);
    }
}
