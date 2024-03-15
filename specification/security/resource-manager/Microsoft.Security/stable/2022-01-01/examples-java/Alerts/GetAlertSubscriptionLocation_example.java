
/**
 * Samples for Alerts GetSubscriptionLevel.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/
     * GetAlertSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get security alert on a subscription from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityAlertOnASubscriptionFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alerts().getSubscriptionLevelWithResponse("westeurope",
            "2518770965529163669_F144EE95-A3E5-42DA-A279-967D115809AA", com.azure.core.util.Context.NONE);
    }
}
