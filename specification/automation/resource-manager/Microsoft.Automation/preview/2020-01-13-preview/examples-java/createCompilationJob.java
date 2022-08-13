import com.azure.resourcemanager.automation.models.DscConfigurationAssociationProperty;

/** Samples for DscCompilationJob Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createCompilationJob.json
     */
    /**
     * Sample code: Create or update a DSC Compilation job.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateADSCCompilationJob(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscCompilationJobs()
            .define("TestCompilationJob")
            .withExistingAutomationAccount("rg", "myAutomationAccount33")
            .withConfiguration(new DscConfigurationAssociationProperty().withName("SetupServer"))
            .create();
    }
}
