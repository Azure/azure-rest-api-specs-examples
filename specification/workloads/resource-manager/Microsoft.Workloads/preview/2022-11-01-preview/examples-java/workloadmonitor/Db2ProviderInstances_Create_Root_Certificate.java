import com.azure.resourcemanager.workloads.models.DB2ProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/Db2ProviderInstances_Create_Root_Certificate.json
     */
    /**
     * Sample code: Create a Db2 provider with Root Certificate.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createADb2ProviderWithRootCertificate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
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
                    .withSslPreference(SslPreference.ROOT_CERTIFICATE))
            .create();
    }
}
