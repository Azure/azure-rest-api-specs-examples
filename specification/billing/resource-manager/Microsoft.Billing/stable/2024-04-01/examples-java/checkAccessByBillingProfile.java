
import com.azure.resourcemanager.billing.models.CheckAccessRequest;
import java.util.Arrays;

/**
 * Samples for BillingPermissions CheckAccessByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByBillingProfile.
     * json
     */
    /**
     * Sample code: CheckAccessByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void checkAccessByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().checkAccessByBillingProfileWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            new CheckAccessRequest().withActions(
                Arrays.asList("Microsoft.Billing/billingAccounts/read", "Microsoft.Subscription/subscriptions/write")),
            com.azure.core.util.Context.NONE);
    }
}
