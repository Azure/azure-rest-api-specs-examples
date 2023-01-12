import com.azure.resourcemanager.billing.fluent.models.CustomerPolicyInner;
import com.azure.resourcemanager.billing.models.ViewCharges;

/** Samples for Policies UpdateCustomer. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateCustomerPolicy.json
     */
    /**
     * Sample code: UpdateCustomer.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updateCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .policies()
            .updateCustomerWithResponse(
                "{billingAccountName}",
                "{customerName}",
                new CustomerPolicyInner().withViewCharges(ViewCharges.NOT_ALLOWED),
                com.azure.core.util.Context.NONE);
    }
}
