
/**
 * Samples for Alerts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/
     * GetAlertsResourceGroup_example.json
     */
    /**
     * Sample code: Get security alerts on a resource group.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityAlertsOnAResourceGroup(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alerts().listByResourceGroup("myRg1", com.azure.core.util.Context.NONE);
    }
}
