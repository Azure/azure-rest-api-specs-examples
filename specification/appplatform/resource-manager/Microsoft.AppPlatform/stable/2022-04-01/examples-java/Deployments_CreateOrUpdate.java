import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.appplatform.fluent.models.DeploymentResourceInner;
import com.azure.resourcemanager.appplatform.models.DeploymentResourceProperties;
import com.azure.resourcemanager.appplatform.models.DeploymentSettings;
import com.azure.resourcemanager.appplatform.models.ResourceRequests;
import com.azure.resourcemanager.appplatform.models.Sku;
import com.azure.resourcemanager.appplatform.models.SourceUploadedUserSourceInfo;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Deployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_CreateOrUpdate.json
     */
    /**
     * Sample code: Deployments_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mydeployment",
                new DeploymentResourceInner()
                    .withProperties(
                        new DeploymentResourceProperties()
                            .withSource(
                                new SourceUploadedUserSourceInfo()
                                    .withVersion("1.0")
                                    .withRelativePath(
                                        "resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc")
                                    .withArtifactSelector("sub-module-1"))
                            .withDeploymentSettings(
                                new DeploymentSettings()
                                    .withResourceRequests(new ResourceRequests().withCpu("1000m").withMemory("3Gi"))
                                    .withEnvironmentVariables(mapOf("env", "test"))
                                    .withAddonConfigs(
                                        mapOf(
                                            "ApplicationConfigurationService",
                                            mapOf(
                                                "patterns",
                                                SerializerFactory
                                                    .createDefaultManagementSerializerAdapter()
                                                    .deserialize(
                                                        "[\"mypattern\"]", Object.class, SerializerEncoding.JSON))))))
                    .withSku(new Sku().withName("S0").withTier("Standard").withCapacity(1)),
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
