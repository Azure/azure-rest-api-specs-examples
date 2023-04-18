import com.azure.resourcemanager.workloads.models.PrometheusHaClusterProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/PrometheusHaClusterProviderInstances_Create_Root_Certificate.json
     */
    /**
     * Sample code: Create a PrometheusHaCluster provider with Root Certificate.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createAPrometheusHaClusterProviderWithRootCertificate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new PrometheusHaClusterProviderInstanceProperties()
                    .withPrometheusUrl("http://192.168.0.0:9090/metrics")
                    .withHostname("hostname")
                    .withSid("sid")
                    .withClusterName("clusterName")
                    .withSslPreference(SslPreference.ROOT_CERTIFICATE))
            .create();
    }
}
