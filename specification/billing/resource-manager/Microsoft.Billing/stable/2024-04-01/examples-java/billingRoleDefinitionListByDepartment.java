
/**
 * Samples for BillingRoleDefinition ListByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionListByDepartment.json
     */
    /**
     * Sample code: BillingRoleDefinitionListByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleDefinitionListByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().listByDepartment("123456", "7368531", com.azure.core.util.Context.NONE);
    }
}
