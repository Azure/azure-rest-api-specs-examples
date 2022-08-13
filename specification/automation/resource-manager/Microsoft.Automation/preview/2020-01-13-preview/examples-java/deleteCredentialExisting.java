import com.azure.core.util.Context;

/** Samples for Credential Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteCredentialExisting.json
     */
    /**
     * Sample code: Delete a credential.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteACredential(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.credentials().deleteWithResponse("rg", "myAutomationAccount20", "myCredential", Context.NONE);
    }
}
