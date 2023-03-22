/** Samples for Automations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/GetAutomationsSubscription_example.json
     */
    /**
     * Sample code: List all security automations of a specified subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllSecurityAutomationsOfASpecifiedSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.automations().list(com.azure.core.util.Context.NONE);
    }
}
