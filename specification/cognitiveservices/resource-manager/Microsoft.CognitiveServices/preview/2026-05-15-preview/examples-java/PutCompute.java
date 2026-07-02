
import com.azure.resourcemanager.cognitiveservices.models.ClusterComputeProperties;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.Pool;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.VmPriority;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Computes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/PutCompute.json
     */
    /**
     * Sample code: PutCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().define("myCompute").withExistingAccount("rgcognitiveservices", "myAccount")
            .withProperties(new ClusterComputeProperties()
                .withPools(Arrays.asList(new Pool().withName("default").withVmPriority(VmPriority.REGULAR)
                    .withInstanceType("Standard_DS3_v2").withNodeCount(2)))
                .withSubnetArmId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/default"))
            .withRegion("eastus").withIdentity(new Identity().withType(ResourceIdentityType.NONE)).create();
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
