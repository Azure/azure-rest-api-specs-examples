
/**
 * Samples for BillingRoleDefinition ListByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionListByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingRoleDefinitionListByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionListByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().listByEnrollmentAccount("123456", "4568789", com.azure.core.util.Context.NONE);
    }
}
