import com.azure.resourcemanager.deploymentmanager.models.DeploymentMode;
import com.azure.resourcemanager.deploymentmanager.models.ServiceUnitArtifacts;
import java.util.HashMap;
import java.util.Map;

/** Samples for ServiceUnits CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate_noartifactsource.json
     */
    /**
     * Sample code: Create service unit using SAS URIs.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createServiceUnitUsingSASURIs(
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
                    .withTemplateUri(
                        "https://mystorageaccount.blob.core.windows.net/myartifactsource/templates/myTopologyUnit.template.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D")
                    .withParametersUri(
                        "https://mystorageaccount.blob.core.windows.net/myartifactsource/parameter/myTopologyUnit.parameters.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"))
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
