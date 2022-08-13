/** Samples for Credential CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateCredential.json
     */
    /**
     * Sample code: Create a credential.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createACredential(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .credentials()
            .define("myCredential")
            .withExistingAutomationAccount("rg", "myAutomationAccount18")
            .withName("myCredential")
            .withUsername("mylingaiah")
            .withPassword("<password>")
            .withDescription("my description goes here")
            .create();
    }
}
