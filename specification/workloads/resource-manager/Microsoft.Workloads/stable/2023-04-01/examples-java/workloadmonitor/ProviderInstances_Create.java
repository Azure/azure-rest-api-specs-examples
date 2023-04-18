import com.azure.resourcemanager.workloads.models.HanaDbProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Create.json
     */
    /**
     * Sample code: Create a SAP monitor Hana provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createASAPMonitorHanaProvider(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new HanaDbProviderInstanceProperties()
                    .withHostname("name")
                    .withDbName("db")
                    .withSqlPort("0000")
                    .withInstanceNumber("00")
                    .withDbUsername("user")
                    .withDbPassword("fakeTokenPlaceholder")
                    .withDbPasswordUri("fakeTokenPlaceholder")
                    .withSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename")
                    .withSslHostnameInCertificate("xyz.domain.com")
                    .withSslPreference(SslPreference.SERVER_CERTIFICATE)
                    .withSapSid("SID"))
            .create();
    }
}
