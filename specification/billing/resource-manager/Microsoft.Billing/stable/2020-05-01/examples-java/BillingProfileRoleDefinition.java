/** Samples for BillingRoleDefinitions GetByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleDefinition.json
     */
    /**
     * Sample code: BillingProfileRoleDefinition.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileRoleDefinition(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleDefinitions()
            .getByBillingProfileWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{billingRoleDefinitionName}",
                com.azure.core.util.Context.NONE);
    }
}
