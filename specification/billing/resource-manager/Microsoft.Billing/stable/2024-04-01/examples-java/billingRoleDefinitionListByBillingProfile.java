
/**
 * Samples for BillingRoleDefinition ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionListByBillingProfile.json
     */
    /**
     * Sample code: BillingRoleDefinitionListByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().listByBillingProfile(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            com.azure.core.util.Context.NONE);
    }
}
