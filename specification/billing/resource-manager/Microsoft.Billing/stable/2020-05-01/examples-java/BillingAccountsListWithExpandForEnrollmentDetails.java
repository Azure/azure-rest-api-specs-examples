import com.azure.core.util.Context;

/** Samples for BillingAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsListWithExpandForEnrollmentDetails.json
     */
    /**
     * Sample code: BillingAccountsListWithExpandForEnrollmentDetails.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsListWithExpandForEnrollmentDetails(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().list("enrollmentDetails,departments,enrollmentAccounts", Context.NONE);
    }
}
