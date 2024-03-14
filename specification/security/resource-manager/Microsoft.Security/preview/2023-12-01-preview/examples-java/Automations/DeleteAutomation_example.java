
/**
 * Samples for Automations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-12-01-preview/examples/Automations/
     * DeleteAutomation_example.json
     */
    /**
     * Sample code: Delete a security automation.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityAutomation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.automations().deleteByResourceGroupWithResponse("myRg", "myAutomationName",
            com.azure.core.util.Context.NONE);
    }
}
