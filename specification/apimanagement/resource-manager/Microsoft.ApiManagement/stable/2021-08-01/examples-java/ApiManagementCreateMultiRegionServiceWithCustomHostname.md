```java
import com.azure.resourcemanager.apimanagement.models.AdditionalLocation;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.ApiVersionConstraint;
import com.azure.resourcemanager.apimanagement.models.HostnameConfiguration;
import com.azure.resourcemanager.apimanagement.models.HostnameType;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
     */
    /**
     * Sample code: ApiManagementCreateMultiRegionServiceWithCustomHostname.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateMultiRegionServiceWithCustomHostname(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .define("apimService1")
            .withRegion("West US")
            .withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(1))
            .withPublisherEmail("apim@autorestsdk.com")
            .withPublisherName("autorestsdk")
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withHostnameConfigurations(
                Arrays
                    .asList(
                        new HostnameConfiguration()
                            .withType(HostnameType.PROXY)
                            .withHostname("gateway1.msitesting.net")
                            .withEncodedCertificate("****** Base 64 Encoded Certificate ************")
                            .withCertificatePassword("Password")
                            .withDefaultSslBinding(true),
                        new HostnameConfiguration()
                            .withType(HostnameType.MANAGEMENT)
                            .withHostname("mgmt.msitesting.net")
                            .withEncodedCertificate("****** Base 64 Encoded Certificate ************")
                            .withCertificatePassword("Password"),
                        new HostnameConfiguration()
                            .withType(HostnameType.PORTAL)
                            .withHostname("portal1.msitesting.net")
                            .withEncodedCertificate("****** Base 64 Encoded Certificate ************")
                            .withCertificatePassword("Password")))
            .withAdditionalLocations(
                Arrays
                    .asList(
                        new AdditionalLocation()
                            .withLocation("East US")
                            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(1))
                            .withDisableGateway(true)))
            .withVirtualNetworkType(VirtualNetworkType.NONE)
            .withApiVersionConstraint(new ApiVersionConstraint().withMinApiVersion("2019-01-01"))
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
