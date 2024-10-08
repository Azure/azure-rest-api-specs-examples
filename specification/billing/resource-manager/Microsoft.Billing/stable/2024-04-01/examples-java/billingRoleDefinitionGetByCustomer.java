
/**
 * Samples for BillingRoleDefinition GetByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByCustomer.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByCustomer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleDefinitionGetByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByCustomerWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "11111111-1111-1111-1111-111111111111", "30000000-aaaa-bbbb-cccc-100000000000",
            com.azure.core.util.Context.NONE);
    }
}
