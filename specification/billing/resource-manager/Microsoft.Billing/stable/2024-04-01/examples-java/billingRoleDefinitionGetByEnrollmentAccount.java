
/**
 * Samples for BillingRoleDefinition GetByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionGetByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByEnrollmentAccountWithResponse("123456", "4568789",
            "50000000-aaaa-bbbb-cccc-100000000000", com.azure.core.util.Context.NONE);
    }
}
