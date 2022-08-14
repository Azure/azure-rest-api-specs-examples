import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Credential;

/** Samples for Credential Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateCredential_patch.json
     */
    /**
     * Sample code: Update a credential.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateACredential(com.azure.resourcemanager.automation.AutomationManager manager) {
        Credential resource =
            manager
                .credentials()
                .getWithResponse("rg", "myAutomationAccount18", "myCredential", Context.NONE)
                .getValue();
        resource
            .update()
            .withName("myCredential")
            .withUsername("mylingaiah")
            .withPassword("<password>")
            .withDescription("my description goes here")
            .apply();
    }
}
