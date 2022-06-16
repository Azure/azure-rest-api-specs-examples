import com.azure.resourcemanager.azurearcdata.models.BasicLoginInformation;
import com.azure.resourcemanager.azurearcdata.models.DataControllerProperties;
import com.azure.resourcemanager.azurearcdata.models.ExtendedLocation;
import com.azure.resourcemanager.azurearcdata.models.ExtendedLocationTypes;
import com.azure.resourcemanager.azurearcdata.models.Infrastructure;
import com.azure.resourcemanager.azurearcdata.models.LogAnalyticsWorkspaceConfig;
import com.azure.resourcemanager.azurearcdata.models.OnPremiseProperty;
import com.azure.resourcemanager.azurearcdata.models.UploadServicePrincipal;
import com.azure.resourcemanager.azurearcdata.models.UploadWatermark;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/** Samples for DataControllers PutDataController. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/CreateOrUpdateDataController.json
     */
    /**
     * Sample code: Create or update a Data Controller.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void createOrUpdateADataController(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager
            .dataControllers()
            .define("testdataController")
            .withRegion("northeurope")
            .withExistingResourceGroup("testrg")
            .withProperties(
                new DataControllerProperties()
                    .withInfrastructure(Infrastructure.ONPREMISES)
                    .withOnPremiseProperty(
                        new OnPremiseProperty()
                            .withId(UUID.fromString("12345678-1234-1234-ab12-1a2b3c4d5e6f"))
                            .withPublicSigningKey("publicOnPremSigningKey"))
                    .withUploadWatermark(
                        new UploadWatermark()
                            .withMetrics(OffsetDateTime.parse("2020-01-01T17:18:19.1234567Z"))
                            .withLogs(OffsetDateTime.parse("2020-01-01T17:18:19.1234567Z"))
                            .withUsages(OffsetDateTime.parse("2020-01-01T17:18:19.1234567Z")))
                    .withBasicLoginInformation(
                        new BasicLoginInformation().withUsername("username").withPassword("********"))
                    .withLogAnalyticsWorkspaceConfig(
                        new LogAnalyticsWorkspaceConfig()
                            .withWorkspaceId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                            .withPrimaryKey("********"))
                    .withUploadServicePrincipal(
                        new UploadServicePrincipal()
                            .withClientId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                            .withTenantId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                            .withAuthority("https://login.microsoftonline.com/")
                            .withClientSecret("********"))
                    .withClusterId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s")
                    .withExtensionId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"))
            .withTags(mapOf("mytag", "myval"))
            .withExtendedLocation(
                new ExtendedLocation()
                    .withName(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation")
                    .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
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
