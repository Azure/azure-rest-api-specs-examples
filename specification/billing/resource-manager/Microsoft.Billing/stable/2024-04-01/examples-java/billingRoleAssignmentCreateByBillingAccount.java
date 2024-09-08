
import com.azure.resourcemanager.billing.models.BillingRoleAssignmentProperties;

/**
 * Samples for BillingRoleAssignments CreateByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentCreateByBillingAccount.json
     */
    /**
     * Sample code: BillingRoleAssignmentCreateByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentCreateByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().createByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30",
            new BillingRoleAssignmentProperties().withPrincipalId("00000000-0000-0000-0000-000000000000")
                .withPrincipalTenantId("076915e7-de10-4323-bb34-a58c904068bb")
                .withRoleDefinitionId(
                    "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingRoleDefinitions/10000000-aaaa-bbbb-cccc-100000000000")
                .withUserEmailAddress("john@contoso.com"),
            com.azure.core.util.Context.NONE);
    }
}
