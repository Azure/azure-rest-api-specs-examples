
import com.azure.resourcemanager.hybridconnectivity.models.AwsCloudProfile;
import com.azure.resourcemanager.hybridconnectivity.models.PublicCloudConnector;
import com.azure.resourcemanager.hybridconnectivity.models.PublicCloudConnectorProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PublicCloudConnectors Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/PublicCloudConnectors_Update.json
     */
    /**
     * Sample code: PublicCloudConnectors_Update.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        publicCloudConnectorsUpdate(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        PublicCloudConnector resource = manager.publicCloudConnectors().getByResourceGroupWithResponse("rgpublicCloud",
            "svtirlbyqpepbzyessjenlueeznhg", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf()).withProperties(new PublicCloudConnectorProperties()
            .withAwsCloudProfile(new AwsCloudProfile().withExcludedAccounts(Arrays.asList("zrbtd")))).apply();
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
