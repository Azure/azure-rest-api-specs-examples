import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_SharedImageGallery.json
     */
    /**
     * Sample code: CreatePool - Custom Image.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolCustomImage(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withVirtualMachineConfiguration(
                        new VirtualMachineConfiguration()
                            .withImageReference(
                                new ImageReference()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1"))
                            .withNodeAgentSkuId("batch.node.ubuntu 18.04")))
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
