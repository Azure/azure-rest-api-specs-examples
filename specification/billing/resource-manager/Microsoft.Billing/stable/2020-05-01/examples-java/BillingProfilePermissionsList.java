/** Samples for BillingPermissions ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilePermissionsList.json
     */
    /**
     * Sample code: BillingProfilePermissionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfilePermissionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingPermissions()
            .listByBillingProfile("{billingAccountName}", "{billingProfileName}", com.azure.core.util.Context.NONE);
    }
}
