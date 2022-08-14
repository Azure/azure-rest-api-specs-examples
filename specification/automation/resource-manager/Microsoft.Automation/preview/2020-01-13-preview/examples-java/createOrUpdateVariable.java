/** Samples for Variable CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateVariable.json
     */
    /**
     * Sample code: Create or update a variable.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateAVariable(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .variables()
            .define("sampleVariable")
            .withExistingAutomationAccount("rg", "sampleAccount9")
            .withName("sampleVariable")
            .withValue("\"ComputerName.domain.com\"")
            .withDescription("my description")
            .withIsEncrypted(false)
            .create();
    }
}
