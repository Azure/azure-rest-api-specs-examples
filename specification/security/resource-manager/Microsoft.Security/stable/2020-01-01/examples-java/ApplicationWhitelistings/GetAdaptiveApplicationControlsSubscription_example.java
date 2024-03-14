
/**
 * Samples for AdaptiveApplicationControls List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/
     * GetAdaptiveApplicationControlsSubscription_example.json
     */
    /**
     * Sample code: Gets a list of application control groups of machines for the subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAListOfApplicationControlGroupsOfMachinesForTheSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.adaptiveApplicationControls().listWithResponse(true, false, com.azure.core.util.Context.NONE);
    }
}
