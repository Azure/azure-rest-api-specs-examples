
import com.azure.resourcemanager.machinelearning.models.AmlCompute;
import com.azure.resourcemanager.machinelearning.models.AmlComputeProperties;
import com.azure.resourcemanager.machinelearning.models.OsType;
import com.azure.resourcemanager.machinelearning.models.RemoteLoginPortPublicAccess;
import com.azure.resourcemanager.machinelearning.models.ScaleSettings;
import com.azure.resourcemanager.machinelearning.models.VirtualMachineImage;
import com.azure.resourcemanager.machinelearning.models.VmPriority;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Compute CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/createOrUpdate/BasicAmlCompute.json
     */
    /**
     * Sample code: Create a AML Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createAAMLCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().define("compute123").withExistingWorkspace("testrg123", "workspaces123").withRegion("eastus")
            .withProperties(new AmlCompute().withProperties(new AmlComputeProperties().withOsType(OsType.WINDOWS)
                .withVmSize("STANDARD_NC6").withVmPriority(VmPriority.DEDICATED)
                .withVirtualMachineImage(new VirtualMachineImage().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myImageGallery/images/myImageDefinition/versions/0.0.1"))
                .withIsolatedNetwork(false)
                .withScaleSettings(new ScaleSettings().withMaxNodeCount(1).withMinNodeCount(0)
                    .withNodeIdleTimeBeforeScaleDown(Duration.parse("PT5M")))
                .withRemoteLoginPortPublicAccess(RemoteLoginPortPublicAccess.NOT_SPECIFIED)
                .withEnableNodePublicIp(true)))
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
