import com.azure.resourcemanager.deploymentmanager.models.DeploymentMode;
import com.azure.resourcemanager.deploymentmanager.models.ServiceUnitArtifacts;
import java.util.HashMap;
import java.util.Map;

/** Samples for ServiceUnits CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate.json
     */
    /**
     * Sample code: Create service unit using relative paths into the artifact source.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createServiceUnitUsingRelativePathsIntoTheArtifactSource(
        com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceUnits()
            .define("myServiceUnit")
            .withRegion("centralus")
            .withExistingService("myResourceGroup", "myTopology", "myService")
            .withTargetResourceGroup("myDeploymentResourceGroup")
            .withDeploymentMode(DeploymentMode.INCREMENTAL)
            .withTags(mapOf())
            .withArtifacts(
                new ServiceUnitArtifacts()
                    .withTemplateArtifactSourceRelativePath("templates/myTopologyUnit.template.json")
                    .withParametersArtifactSourceRelativePath("parameter/myTopologyUnit.parameters.json"))
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
