import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceIdentity;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.ApiVersionConstraint;
import com.azure.resourcemanager.apimanagement.models.ApimIdentityType;
import com.azure.resourcemanager.apimanagement.models.HostnameConfiguration;
import com.azure.resourcemanager.apimanagement.models.HostnameType;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import com.azure.resourcemanager.apimanagement.models.UserIdentityProperties;
import com.azure.resourcemanager.apimanagement.models.VirtualNetworkType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceWithCustomHostnameKeyVault.json
     */
    /**
     * Sample code: ApiManagementCreateServiceWithCustomHostnameKeyVault.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateServiceWithCustomHostnameKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .define("apimService1")
            .withRegion("North Europe")
            .withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.PREMIUM).withCapacity(1))
            .withPublisherEmail("apim@autorestsdk.com")
            .withPublisherName("autorestsdk")
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withIdentity(
                new ApiManagementServiceIdentity()
                    .withType(ApimIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserIdentityProperties())))
            .withHostnameConfigurations(
                Arrays
                    .asList(
                        new HostnameConfiguration()
                            .withType(HostnameType.PROXY)
                            .withHostname("gateway1.msitesting.net")
                            .withKeyVaultId("fakeTokenPlaceholder")
                            .withIdentityClientId("329419bc-adec-4dce-9568-25a6d486e468")
                            .withDefaultSslBinding(true),
                        new HostnameConfiguration()
                            .withType(HostnameType.MANAGEMENT)
                            .withHostname("mgmt.msitesting.net")
                            .withKeyVaultId("fakeTokenPlaceholder")
                            .withIdentityClientId("329419bc-adec-4dce-9568-25a6d486e468"),
                        new HostnameConfiguration()
                            .withType(HostnameType.PORTAL)
                            .withHostname("portal1.msitesting.net")
                            .withKeyVaultId("fakeTokenPlaceholder")
                            .withIdentityClientId("329419bc-adec-4dce-9568-25a6d486e468")))
            .withVirtualNetworkType(VirtualNetworkType.NONE)
            .withApiVersionConstraint(new ApiVersionConstraint().withMinApiVersion("2019-01-01"))
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
