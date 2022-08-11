import com.azure.core.util.Context;
import com.azure.resourcemanager.automanage.fluent.models.ConfigurationProfileAssignmentInner;
import com.azure.resourcemanager.automanage.models.ConfigurationProfileAssignmentProperties;

/** Samples for ConfigurationProfileHcrpAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileHCRPAssignment.json
     */
    /**
     * Sample code: Create or update HCRP configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void createOrUpdateHCRPConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileHcrpAssignments()
            .createOrUpdateWithResponse(
                "myResourceGroupName",
                "myMachineName",
                "default",
                new ConfigurationProfileAssignmentInner()
                    .withProperties(
                        new ConfigurationProfileAssignmentProperties()
                            .withConfigurationProfile(
                                "/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction")),
                Context.NONE);
    }
}
