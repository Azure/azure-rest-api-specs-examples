Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
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
                Context.NONE);
    }
}
```
