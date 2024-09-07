
/**
 * Samples for EnrollmentAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountGet.json
     */
    /**
     * Sample code: EnrollmentAccountGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().getWithResponse("6564892", "257698", com.azure.core.util.Context.NONE);
    }
}
