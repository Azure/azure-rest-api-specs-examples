/** Samples for EnrollmentAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsList.json
     */
    /**
     * Sample code: EnrollmentAccountsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().list(com.azure.core.util.Context.NONE);
    }
}
