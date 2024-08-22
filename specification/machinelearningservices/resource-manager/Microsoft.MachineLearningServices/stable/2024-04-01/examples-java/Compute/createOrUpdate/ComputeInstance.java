
import com.azure.resourcemanager.machinelearning.models.ApplicationSharingPolicy;
import com.azure.resourcemanager.machinelearning.models.AssignedUser;
import com.azure.resourcemanager.machinelearning.models.ComputeInstance;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceAuthorizationType;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceProperties;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceSshSettings;
import com.azure.resourcemanager.machinelearning.models.CustomService;
import com.azure.resourcemanager.machinelearning.models.Docker;
import com.azure.resourcemanager.machinelearning.models.Endpoint;
import com.azure.resourcemanager.machinelearning.models.EnvironmentVariable;
import com.azure.resourcemanager.machinelearning.models.EnvironmentVariableType;
import com.azure.resourcemanager.machinelearning.models.Image;
import com.azure.resourcemanager.machinelearning.models.ImageType;
import com.azure.resourcemanager.machinelearning.models.PersonalComputeInstanceSettings;
import com.azure.resourcemanager.machinelearning.models.Protocol;
import com.azure.resourcemanager.machinelearning.models.ResourceId;
import com.azure.resourcemanager.machinelearning.models.SshPublicAccess;
import com.azure.resourcemanager.machinelearning.models.VolumeDefinition;
import com.azure.resourcemanager.machinelearning.models.VolumeDefinitionType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Compute CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/createOrUpdate/ComputeInstance.json
     */
    /**
     * Sample code: Create an ComputeInstance Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        createAnComputeInstanceCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().define("compute123").withExistingWorkspace("testrg123", "workspaces123").withRegion("eastus")
            .withProperties(new ComputeInstance().withProperties(new ComputeInstanceProperties()
                .withVmSize("STANDARD_NC6").withSubnet(new ResourceId().withId("test-subnet-resource-id"))
                .withApplicationSharingPolicy(ApplicationSharingPolicy.PERSONAL)
                .withSshSettings(new ComputeInstanceSshSettings().withSshPublicAccess(SshPublicAccess.DISABLED))
                .withCustomServices(Arrays.asList(new CustomService().withName("rstudio")
                    .withImage(new Image().withType(ImageType.DOCKER)
                        .withReference("ghcr.io/azure/rocker-rstudio-ml-verse:latest")
                        .withAdditionalProperties(mapOf()))
                    .withEnvironmentVariables(mapOf("test_variable",
                        new EnvironmentVariable().withType(EnvironmentVariableType.LOCAL).withValue("test_value")
                            .withAdditionalProperties(mapOf())))
                    .withDocker(new Docker().withPrivileged(true).withAdditionalProperties(mapOf()))
                    .withEndpoints(Arrays.asList(new Endpoint().withProtocol(Protocol.HTTP).withName("connect")
                        .withTarget(8787).withPublished(8787)))
                    .withVolumes(
                        Arrays.asList(new VolumeDefinition().withType(VolumeDefinitionType.BIND).withReadOnly(false)
                            .withSource("/home/azureuser/cloudfiles").withTarget("/home/azureuser/cloudfiles")))
                    .withAdditionalProperties(mapOf())))
                .withComputeInstanceAuthorizationType(ComputeInstanceAuthorizationType.PERSONAL)
                .withPersonalComputeInstanceSettings(new PersonalComputeInstanceSettings()
                    .withAssignedUser(new AssignedUser().withObjectId("00000000-0000-0000-0000-000000000000")
                        .withTenantId("00000000-0000-0000-0000-000000000000")))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
