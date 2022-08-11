import com.azure.resourcemanager.automanage.models.ConfigurationProfileAssignmentProperties;

/** Samples for ConfigurationProfileAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileAssignment.json
     */
    /**
     * Sample code: Create or update configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void createOrUpdateConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileAssignments()
            .define("default")
            .withExistingVirtualMachine("myResourceGroupName", "myVMName")
            .withProperties(
                new ConfigurationProfileAssignmentProperties()
                    .withConfigurationProfile(
                        "/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"))
            .create();
    }
}
