import com.azure.resourcemanager.deploymentmanager.models.Identity;
import com.azure.resourcemanager.deploymentmanager.models.PrePostStep;
import com.azure.resourcemanager.deploymentmanager.models.StepGroup;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Rollouts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_createorupdate.json
     */
    /**
     * Sample code: Create or update rollout.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createOrUpdateRollout(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .rollouts()
            .define("myRollout")
            .withRegion("centralus")
            .withExistingResourceGroup("myResourceGroup")
            .withIdentity(
                new Identity()
                    .withType("userAssigned")
                    .withIdentityIds(
                        Arrays
                            .asList(
                                "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity")))
            .withBuildVersion("1.0.0.1")
            .withTargetServiceTopologyId(
                "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/serviceTopologies/myTopology")
            .withStepGroups(
                Arrays
                    .asList(
                        new StepGroup()
                            .withName("FirstRegion")
                            .withPreDeploymentSteps(
                                Arrays
                                    .asList(
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/preDeployStep1"),
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/preDeployStep2")))
                            .withDeploymentTargetId(
                                "Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit1'")
                            .withPostDeploymentSteps(
                                Arrays
                                    .asList(
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/postDeployStep1"))),
                        new StepGroup()
                            .withName("SecondRegion")
                            .withDependsOnStepGroups(Arrays.asList("FirstRegion"))
                            .withPreDeploymentSteps(
                                Arrays
                                    .asList(
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/preDeployStep3"),
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/preDeployStep4")))
                            .withDeploymentTargetId(
                                "Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit2'")
                            .withPostDeploymentSteps(
                                Arrays
                                    .asList(
                                        new PrePostStep()
                                            .withStepId("Microsoft.DeploymentManager/steps/postDeployStep5")))))
            .withTags(mapOf())
            .withArtifactSourceId(
                "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/artifactSources/myArtifactSource")
            .create();
    }

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
