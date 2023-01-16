import com.azure.resourcemanager.deploymentmanager.models.WaitStepAttributes;
import com.azure.resourcemanager.deploymentmanager.models.WaitStepProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Steps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_wait_createorupdate.json
     */
    /**
     * Sample code: Create wait step.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createWaitStep(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .steps()
            .define("waitStep")
            .withRegion("centralus")
            .withExistingResourceGroup("myResourceGroup")
            .withProperties(new WaitStepProperties().withAttributes(new WaitStepAttributes().withDuration("PT20M")))
            .withTags(mapOf())
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
