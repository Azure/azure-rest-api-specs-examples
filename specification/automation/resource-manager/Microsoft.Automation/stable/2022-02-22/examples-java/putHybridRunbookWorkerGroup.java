import com.azure.resourcemanager.automation.models.RunAsCredentialAssociationProperty;

/** Samples for HybridRunbookWorkerGroup Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/putHybridRunbookWorkerGroup.json
     */
    /**
     * Sample code: Create a hybrid worker group.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createAHybridWorkerGroup(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .hybridRunbookWorkerGroups()
            .define("TestHybridGroup")
            .withExistingAutomationAccount("rg", "testaccount")
            .withCredential(new RunAsCredentialAssociationProperty().withName("myRunAsCredentialName"))
            .create();
    }
}
