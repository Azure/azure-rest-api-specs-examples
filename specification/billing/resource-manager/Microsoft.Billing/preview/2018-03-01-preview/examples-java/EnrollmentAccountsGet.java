import com.azure.core.util.Context;

/** Samples for EnrollmentAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsGet.json
     */
    /**
     * Sample code: EnrollmentAccountsGet.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountsGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().getWithResponse("e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556", Context.NONE);
    }
}
