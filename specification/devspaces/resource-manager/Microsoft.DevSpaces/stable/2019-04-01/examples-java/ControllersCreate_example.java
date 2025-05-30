
import com.azure.resourcemanager.devspaces.models.Sku;
import com.azure.resourcemanager.devspaces.models.SkuName;
import com.azure.resourcemanager.devspaces.models.SkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Controllers Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersCreate_example
     * .json
     */
    /**
     * Sample code: ControllersCreate.
     * 
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersCreate(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager.controllers().define("myControllerResource").withRegion("eastus")
            .withExistingResourceGroup("myResourceGroup")
            .withSku(new Sku().withName(SkuName.S1).withTier(SkuTier.STANDARD))
            .withTargetContainerHostResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster")
            .withTargetContainerHostCredentialsBase64("QmFzZTY0IEVuY29kZWQgVmFsdWUK").withTags(mapOf()).create();
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
