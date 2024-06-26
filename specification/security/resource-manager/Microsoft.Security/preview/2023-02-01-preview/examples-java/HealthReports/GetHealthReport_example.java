/** Samples for HealthReportOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-02-01-preview/examples/HealthReports/GetHealthReport_example.json
     */
    /**
     * Sample code: Get health report of resource.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getHealthReportOfResource(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .healthReportOperations()
            .getWithResponse(
                "subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings",
                "909c629a-bf39-4521-8e4f-10b443a0bc02",
                com.azure.core.util.Context.NONE);
    }
}
