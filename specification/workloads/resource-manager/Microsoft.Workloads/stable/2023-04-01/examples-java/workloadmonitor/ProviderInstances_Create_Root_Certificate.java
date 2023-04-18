import com.azure.resourcemanager.workloads.models.HanaDbProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Create_Root_Certificate.json
     */
    /**
     * Sample code: Create a SAP monitor Hana provider with Root Certificate.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createASAPMonitorHanaProviderWithRootCertificate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
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
                    .withSslHostnameInCertificate("xyz.domain.com")
                    .withSslPreference(SslPreference.ROOT_CERTIFICATE)
                    .withSapSid("SID"))
            .create();
    }
}
