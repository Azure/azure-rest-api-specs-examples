import com.azure.resourcemanager.devtestlabs.models.EnvironmentDeploymentProperties;
import java.util.Arrays;

/** Samples for Environments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_CreateOrUpdate.json
     */
    /**
     * Sample code: Environments_CreateOrUpdate.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void environmentsCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .environments()
            .define("{environmentName}")
            .withRegion((String) null)
            .withExistingUser("resourceGroupName", "{labName}", "@me")
            .withDeploymentProperties(
                new EnvironmentDeploymentProperties()
                    .withArmTemplateId(
                        "/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/armTemplates/{armTemplateName}")
                    .withParameters(Arrays.asList()))
            .create();
    }
}
