
import com.azure.resourcemanager.cognitiveservices.models.AccessKeyAuthTypeConnectionProperties;
import com.azure.resourcemanager.cognitiveservices.models.ConnectionAccessKey;
import com.azure.resourcemanager.cognitiveservices.models.ConnectionCategory;
import com.azure.resourcemanager.cognitiveservices.models.ConnectionUpdateContent;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ProjectConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ProjectConnection/update.json
     */
    /**
     * Sample code: UpdateProjectConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        updateProjectConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectConnections().updateWithResponse("test-rg", "account-1", "project-1", "connection-1",
            new ConnectionUpdateContent().withProperties(new AccessKeyAuthTypeConnectionProperties()
                .withCategory(ConnectionCategory.ADLSGEN2).withExpiryTime(OffsetDateTime.parse("2020-01-01T00:00:00Z"))
                .withMetadata(mapOf()).withTarget("some_string").withCredentials(new ConnectionAccessKey()
                    .withAccessKeyId("fakeTokenPlaceholder").withSecretAccessKey("fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
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
