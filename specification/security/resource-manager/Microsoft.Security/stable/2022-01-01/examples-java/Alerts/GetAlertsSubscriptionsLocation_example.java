
/**
 * Samples for Alerts ListSubscriptionLevelByRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/
     * GetAlertsSubscriptionsLocation_example.json
     */
    /**
     * Sample code: Get security alerts on a subscription from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityAlertsOnASubscriptionFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alerts().listSubscriptionLevelByRegion("westeurope", com.azure.core.util.Context.NONE);
    }
}
