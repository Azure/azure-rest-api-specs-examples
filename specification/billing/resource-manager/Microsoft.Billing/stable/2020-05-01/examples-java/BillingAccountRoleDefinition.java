import com.azure.core.util.Context;

/** Samples for BillingRoleDefinitions GetByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleDefinition.json
     */
    /**
     * Sample code: BillingAccountRoleDefinition.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountRoleDefinition(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleDefinitions()
            .getByBillingAccountWithResponse("{billingAccountName}", "{billingRoleDefinitionName}", Context.NONE);
    }
}
