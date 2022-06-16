import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.AppResourceInner;
import com.azure.resourcemanager.appplatform.models.AppResourceProperties;
import com.azure.resourcemanager.appplatform.models.LoadedCertificate;
import com.azure.resourcemanager.appplatform.models.ManagedIdentityProperties;
import com.azure.resourcemanager.appplatform.models.ManagedIdentityType;
import com.azure.resourcemanager.appplatform.models.PersistentDisk;
import com.azure.resourcemanager.appplatform.models.TemporaryDisk;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Apps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_CreateOrUpdate.json
     */
    /**
     * Sample code: Apps_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "myapp",
                new AppResourceInner()
                    .withProperties(
                        new AppResourceProperties()
                            .withPublicProperty(true)
                            .withAddonConfigs(
                                mapOf(
                                    "ApplicationConfigurationService",
                                    mapOf(
                                        "resourceId",
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/configurationServices/myacs"),
                                    "ServiceRegistry",
                                    mapOf(
                                        "resourceId",
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/serviceRegistries/myServiceRegistry")))
                            .withFqdn("myapp.mydomain.com")
                            .withHttpsOnly(false)
                            .withTemporaryDisk(new TemporaryDisk().withSizeInGB(2).withMountPath("/mytemporarydisk"))
                            .withPersistentDisk(new PersistentDisk().withSizeInGB(2).withMountPath("/mypersistentdisk"))
                            .withEnableEndToEndTls(false)
                            .withLoadedCertificates(
                                Arrays
                                    .asList(
                                        new LoadedCertificate()
                                            .withResourceId(
                                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1")
                                            .withLoadTrustStore(false),
                                        new LoadedCertificate()
                                            .withResourceId(
                                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2")
                                            .withLoadTrustStore(true))))
                    .withIdentity(new ManagedIdentityProperties().withType(ManagedIdentityType.SYSTEM_ASSIGNED))
                    .withLocation("eastus"),
                Context.NONE);
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
