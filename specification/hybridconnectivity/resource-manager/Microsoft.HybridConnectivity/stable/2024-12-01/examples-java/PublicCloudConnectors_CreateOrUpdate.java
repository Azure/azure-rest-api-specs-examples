
import com.azure.resourcemanager.hybridconnectivity.models.AwsCloudProfile;
import com.azure.resourcemanager.hybridconnectivity.models.HostType;
import com.azure.resourcemanager.hybridconnectivity.models.PublicCloudConnectorProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PublicCloudConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/PublicCloudConnectors_CreateOrUpdate.json
     */
    /**
     * Sample code: PublicCloudConnectors_CreateOrUpdate.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void publicCloudConnectorsCreateOrUpdate(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.publicCloudConnectors().define("advjwoakdusalamomg").withRegion("jpiglusfxynfcewcjwvvnn")
            .withExistingResourceGroup("rgpublicCloud").withTags(mapOf())
            .withProperties(new PublicCloudConnectorProperties()
                .withAwsCloudProfile(new AwsCloudProfile().withAccountId("snbnuxckevyqpm")
                    .withExcludedAccounts(Arrays.asList("rwgqpukglvbqmogqcliqolucp")).withIsOrganizationalAccount(true))
                .withHostType(HostType.AWS))
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
