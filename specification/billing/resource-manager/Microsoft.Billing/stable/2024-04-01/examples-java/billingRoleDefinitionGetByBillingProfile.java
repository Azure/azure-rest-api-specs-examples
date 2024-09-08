
/**
 * Samples for BillingRoleDefinition GetByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByBillingProfile.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionGetByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByBillingProfileWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "40000000-aaaa-bbbb-cccc-100000000000", com.azure.core.util.Context.NONE);
    }
}
