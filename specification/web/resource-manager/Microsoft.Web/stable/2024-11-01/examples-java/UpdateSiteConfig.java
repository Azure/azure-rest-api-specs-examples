
import com.azure.resourcemanager.appservice.fluent.models.SiteConfigResourceInner;
import com.azure.resourcemanager.appservice.models.FtpsState;
import com.azure.resourcemanager.appservice.models.ManagedPipelineMode;
import com.azure.resourcemanager.appservice.models.SiteLoadBalancing;
import com.azure.resourcemanager.appservice.models.SupportedTlsVersions;
import com.azure.resourcemanager.appservice.models.VirtualApplication;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps CreateOrUpdateConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateSiteConfig.json
     */
    /**
     * Sample code: Update Site Config.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateSiteConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().createOrUpdateConfiguration("testrg123", "sitef6141",
            new SiteConfigResourceInner().withNumberOfWorkers(1)
                .withDefaultDocuments(Arrays.asList("Default.htm", "Default.html", "Default.asp", "index.htm",
                    "index.html", "iisstart.htm", "default.aspx", "index.php", "hostingstart.html"))
                .withNetFrameworkVersion("v4.0").withPhpVersion("5.6").withPythonVersion("").withNodeVersion("")
                .withPowerShellVersion("").withLinuxFxVersion("").withRequestTracingEnabled(false)
                .withRemoteDebuggingEnabled(false).withHttpLoggingEnabled(false).withAcrUseManagedIdentityCreds(false)
                .withLogsDirectorySizeLimit(35).withDetailedErrorLoggingEnabled(false).withUse32BitWorkerProcess(true)
                .withWebSocketsEnabled(false).withAlwaysOn(false).withAppCommandLine("")
                .withManagedPipelineMode(ManagedPipelineMode.INTEGRATED)
                .withVirtualApplications(Arrays.asList(new VirtualApplication().withVirtualPath("/")
                    .withPhysicalPath("site\\wwwroot").withPreloadEnabled(false)))
                .withLoadBalancing(SiteLoadBalancing.LEAST_REQUESTS).withAutoHealEnabled(false).withVnetName("")
                .withVnetRouteAllEnabled(false).withVnetPrivatePortsCount(0).withHttp20Enabled(false)
                .withMinTlsVersion(SupportedTlsVersions.ONE_TWO).withScmMinTlsVersion(SupportedTlsVersions.ONE_TWO)
                .withFtpsState(FtpsState.ALL_ALLOWED).withFunctionAppScaleLimit(0)
                .withFunctionsRuntimeScaleMonitoringEnabled(false).withMinimumElasticInstanceCount(0)
                .withAzureStorageAccounts(mapOf()),
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
