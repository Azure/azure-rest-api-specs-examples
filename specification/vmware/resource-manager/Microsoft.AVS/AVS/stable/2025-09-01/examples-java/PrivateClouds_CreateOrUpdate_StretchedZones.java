
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.Sku;
import com.azure.resourcemanager.avs.models.Vcf5License;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateClouds CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_CreateOrUpdate_StretchedZones.json
     */
    /**
     * Sample code: PrivateClouds_CreateOrUpdate_StretchedZones.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsCreateOrUpdateStretchedZones(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().define("cloud1").withRegion("eastus2").withExistingResourceGroup("group1")
            .withSku(new Sku().withName("AV36")).withTags(mapOf()).withZones(Arrays.asList("1", "2"))
            .withManagementCluster(new ManagementCluster().withClusterSize(4)).withNetworkBlock("192.168.48.0/22")
            .withVcfLicense(new Vcf5License().withLicenseKey("fakeTokenPlaceholder").withCores(16)
                .withEndDate(OffsetDateTime.parse("2025-12-31T23:59:59Z")).withBroadcomSiteId("123456")
                .withBroadcomContractNumber("123456"))
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
