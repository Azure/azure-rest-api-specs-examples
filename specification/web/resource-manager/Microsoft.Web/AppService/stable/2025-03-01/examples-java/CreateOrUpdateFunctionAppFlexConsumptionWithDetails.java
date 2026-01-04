
import com.azure.resourcemanager.appservice.fluent.models.SiteConfigInner;
import com.azure.resourcemanager.appservice.fluent.models.SiteInner;
import com.azure.resourcemanager.appservice.models.AuthenticationType;
import com.azure.resourcemanager.appservice.models.FunctionAppConfig;
import com.azure.resourcemanager.appservice.models.FunctionsAlwaysReadyConfig;
import com.azure.resourcemanager.appservice.models.FunctionsDeployment;
import com.azure.resourcemanager.appservice.models.FunctionsDeploymentStorage;
import com.azure.resourcemanager.appservice.models.FunctionsDeploymentStorageAuthentication;
import com.azure.resourcemanager.appservice.models.FunctionsDeploymentStorageType;
import com.azure.resourcemanager.appservice.models.FunctionsRuntime;
import com.azure.resourcemanager.appservice.models.FunctionsScaleAndConcurrency;
import com.azure.resourcemanager.appservice.models.FunctionsScaleAndConcurrencyTriggers;
import com.azure.resourcemanager.appservice.models.FunctionsScaleAndConcurrencyTriggersHttp;
import com.azure.resourcemanager.appservice.models.NameValuePair;
import com.azure.resourcemanager.appservice.models.RuntimeName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * CreateOrUpdateFunctionAppFlexConsumptionWithDetails.json
     */
    /**
     * Sample code: Create or Update Flex Consumption function app with details.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateFlexConsumptionFunctionAppWithDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().createOrUpdate("testrg123", "sitef6141", new SiteInner()
            .withLocation("East US").withKind("functionapp,linux")
            .withSiteConfig(new SiteConfigInner().withAppSettings(Arrays.asList(
                new NameValuePair().withName("AzureWebJobsStorage").withValue(
                    "DefaultEndpointsProtocol=https;AccountName=StorageAccountName;AccountKey=Sanitized;EndpointSuffix=core.windows.net"),
                new NameValuePair().withName("APPLICATIONINSIGHTS_CONNECTION_STRING")
                    .withValue("InstrumentationKey=Sanitized;IngestionEndpoint=Sanitized;LiveEndpoint=Sanitized"))))
            .withFunctionAppConfig(new FunctionAppConfig()
                .withDeployment(new FunctionsDeployment().withStorage(
                    new FunctionsDeploymentStorage().withType(FunctionsDeploymentStorageType.BLOB_CONTAINER)
                        .withValue("https://storageAccountName.blob.core.windows.net/containername")
                        .withAuthentication(new FunctionsDeploymentStorageAuthentication()
                            .withType(AuthenticationType.STORAGE_ACCOUNT_CONNECTION_STRING)
                            .withStorageAccountConnectionStringName("TheAppSettingName"))))
                .withRuntime(new FunctionsRuntime().withName(RuntimeName.PYTHON).withVersion("3.11"))
                .withScaleAndConcurrency(new FunctionsScaleAndConcurrency()
                    .withAlwaysReady(
                        Arrays.asList(new FunctionsAlwaysReadyConfig().withName("http").withInstanceCount(2)))
                    .withMaximumInstanceCount(100).withInstanceMemoryMB(2048)
                    .withTriggers(new FunctionsScaleAndConcurrencyTriggers()
                        .withHttp(new FunctionsScaleAndConcurrencyTriggersHttp().withPerInstanceConcurrency(16))))),
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
