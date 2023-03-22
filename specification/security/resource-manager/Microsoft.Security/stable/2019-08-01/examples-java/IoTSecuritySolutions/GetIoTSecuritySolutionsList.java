/** Samples for IotSecuritySolution List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsList.json
     */
    /**
     * Sample code: List IoT Security solutions by subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listIoTSecuritySolutionsBySubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutions().list(null, com.azure.core.util.Context.NONE);
    }
}
