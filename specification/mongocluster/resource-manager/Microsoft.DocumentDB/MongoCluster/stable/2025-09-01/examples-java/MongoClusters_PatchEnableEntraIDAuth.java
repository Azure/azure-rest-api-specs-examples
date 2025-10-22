
import com.azure.resourcemanager.mongocluster.models.AuthConfigProperties;
import com.azure.resourcemanager.mongocluster.models.AuthenticationMode;
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/MongoClusters_PatchEnableEntraIDAuth.json
     */
    /**
     * Sample code: Updates the allowed authentication modes to include Microsoft Entra ID authentication.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void updatesTheAllowedAuthenticationModesToIncludeMicrosoftEntraIDAuthentication(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties().withAuthConfig(new AuthConfigProperties()
            .withAllowedModes(Arrays.asList(AuthenticationMode.NATIVE_AUTH, AuthenticationMode.MICROSOFT_ENTRA_ID))))
            .apply();
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
