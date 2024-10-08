
/**
 * Samples for BillingRoleDefinition GetByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByDepartment.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleDefinitionGetByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByDepartmentWithResponse("123456", "7368531",
            "50000000-aaaa-bbbb-cccc-100000000000", com.azure.core.util.Context.NONE);
    }
}
