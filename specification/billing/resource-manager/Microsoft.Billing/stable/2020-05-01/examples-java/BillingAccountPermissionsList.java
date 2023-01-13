/** Samples for BillingPermissions ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountPermissionsList.json
     */
    /**
     * Sample code: BillingAccountPermissionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountPermissionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().listByBillingAccount("{billingAccountName}", com.azure.core.util.Context.NONE);
    }
}
