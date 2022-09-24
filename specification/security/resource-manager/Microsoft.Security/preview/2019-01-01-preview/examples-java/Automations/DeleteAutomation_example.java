import com.azure.core.util.Context;

/** Samples for Automations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/DeleteAutomation_example.json
     */
    /**
     * Sample code: Delete a security automation.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityAutomation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.automations().deleteWithResponse("myRg", "myAutomationName", Context.NONE);
    }
}
