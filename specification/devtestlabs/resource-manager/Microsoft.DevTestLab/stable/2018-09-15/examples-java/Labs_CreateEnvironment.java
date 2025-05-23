
import com.azure.resourcemanager.devtestlabs.models.GalleryImageReference;
import com.azure.resourcemanager.devtestlabs.models.LabVirtualMachineCreationParameter;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Labs CreateEnvironment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateEnvironment
     * .json
     */
    /**
     * Sample code: Labs_CreateEnvironment.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsCreateEnvironment(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().createEnvironment("resourceGroupName", "{labName}", new LabVirtualMachineCreationParameter()
            .withName("{vmName}").withLocation("{location}").withTags(mapOf("tagName1", "tagValue1"))
            .withSize("Standard_A2_v2").withUsername("{userName}").withPassword("fakeTokenPlaceholder")
            .withLabSubnetName("{virtualnetwork-subnet-name}")
            .withLabVirtualNetworkId(
                "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}")
            .withDisallowPublicIpAddress(true)
            .withGalleryImageReference(new GalleryImageReference().withOffer("UbuntuServer").withPublisher("Canonical")
                .withSku("16.04-LTS").withOsType("Linux").withVersion("Latest"))
            .withAllowClaim(true).withStorageType("Standard"), com.azure.core.util.Context.NONE);
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
