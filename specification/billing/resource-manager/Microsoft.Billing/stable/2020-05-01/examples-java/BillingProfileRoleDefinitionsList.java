import com.azure.core.util.Context;

/** Samples for BillingRoleDefinitions ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleDefinitionsList.json
     */
    /**
     * Sample code: BillingProfileRoleDefinitionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileRoleDefinitionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleDefinitions()
            .listByBillingProfile("{billingAccountName}", "{billingProfileName}", Context.NONE);
    }
}
