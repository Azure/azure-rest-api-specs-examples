import com.azure.resourcemanager.workloads.models.DB2ProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/Db2ProviderInstances_Create.json
     */
    /**
     * Sample code: Create a Db2 provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createADb2Provider(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new DB2ProviderInstanceProperties()
                    .withHostname("hostname")
                    .withDbName("dbName")
                    .withDbPort("dbPort")
                    .withDbUsername("username")
                    .withDbPassword("fakeTokenPlaceholder")
                    .withDbPasswordUri("fakeTokenPlaceholder")
                    .withSapSid("SID")
                    .withSslPreference(SslPreference.SERVER_CERTIFICATE)
                    .withSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename"))
            .create();
    }
}
