
/**
 * Samples for BillingRoleDefinition GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByBillingAccount.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionGetByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByBillingAccountWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "50000000-aaaa-bbbb-cccc-100000000000", com.azure.core.util.Context.NONE);
    }
}
