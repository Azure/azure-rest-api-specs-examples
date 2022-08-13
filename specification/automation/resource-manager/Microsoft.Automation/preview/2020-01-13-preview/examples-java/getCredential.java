import com.azure.core.util.Context;

/** Samples for Credential Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCredential.json
     */
    /**
     * Sample code: Get a credential.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getACredential(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.credentials().getWithResponse("rg", "myAutomationAccount18", "myCredential", Context.NONE);
    }
}
