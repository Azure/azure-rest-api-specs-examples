
/**
 * Samples for HealthReports List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-05-01-preview/examples/HealthReports/
     * ListHealthReports_example.json
     */
    /**
     * Sample code: List health reports.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listHealthReports(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.healthReports().list("subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2",
            com.azure.core.util.Context.NONE);
    }
}
