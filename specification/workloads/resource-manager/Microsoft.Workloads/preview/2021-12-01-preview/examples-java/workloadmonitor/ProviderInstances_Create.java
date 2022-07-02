import com.azure.resourcemanager.workloads.models.HanaDbProviderInstanceProperties;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/ProviderInstances_Create.json
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
                    .withDbPassword("****")
                    .withDbPasswordUri("")
                    .withDbSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename")
                    .withSslHostnameInCertificate("xyz.domain.com"))
            .create();
    }
}
