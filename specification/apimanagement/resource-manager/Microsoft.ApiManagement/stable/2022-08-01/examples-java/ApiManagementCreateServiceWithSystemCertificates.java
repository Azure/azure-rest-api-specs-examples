import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.CertificateConfiguration;
import com.azure.resourcemanager.apimanagement.models.CertificateConfigurationStoreName;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceWithSystemCertificates.json
     */
    /**
     * Sample code: ApiManagementCreateServiceWithSystemCertificates.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateServiceWithSystemCertificates(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .define("apimService1")
            .withRegion("Central US")
            .withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.BASIC).withCapacity(1))
            .withPublisherEmail("apim@autorestsdk.com")
            .withPublisherName("autorestsdk")
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withCertificates(
                Arrays
                    .asList(
                        new CertificateConfiguration()
                            .withEncodedCertificate("fakeTokenPlaceholder")
                            .withCertificatePassword("fakeTokenPlaceholder")
                            .withStoreName(CertificateConfigurationStoreName.CERTIFICATE_AUTHORITY)))
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
