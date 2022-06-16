import com.azure.core.util.Context;

/** Samples for BillingRoleDefinitions ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleDefinitionsList.json
     */
    /**
     * Sample code: BillingAccountRoleDefinitionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountRoleDefinitionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().listByBillingAccount("{billingAccountName}", Context.NONE);
    }
}
