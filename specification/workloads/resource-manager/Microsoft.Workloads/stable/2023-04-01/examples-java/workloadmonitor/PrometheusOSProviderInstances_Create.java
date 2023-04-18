import com.azure.resourcemanager.workloads.models.PrometheusOSProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/PrometheusOSProviderInstances_Create.json
     */
    /**
     * Sample code: Create a OS provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createAOSProvider(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new PrometheusOSProviderInstanceProperties()
                    .withPrometheusUrl("http://192.168.0.0:9090/metrics")
                    .withSslPreference(SslPreference.SERVER_CERTIFICATE)
                    .withSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename")
                    .withSapSid("SID"))
            .create();
    }
}
