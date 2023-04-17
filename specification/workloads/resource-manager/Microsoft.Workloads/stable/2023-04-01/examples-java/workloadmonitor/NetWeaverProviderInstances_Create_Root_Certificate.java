import com.azure.resourcemanager.workloads.models.SapNetWeaverProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;
import java.util.Arrays;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/NetWeaverProviderInstances_Create_Root_Certificate.json
     */
    /**
     * Sample code: Create a SAP monitor NetWeaver provider with Root Certificate.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createASAPMonitorNetWeaverProviderWithRootCertificate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new SapNetWeaverProviderInstanceProperties()
                    .withSapSid("SID")
                    .withSapHostname("name")
                    .withSapInstanceNr("00")
                    .withSapHostFileEntries(Arrays.asList("127.0.0.1 name fqdn"))
                    .withSapUsername("username")
                    .withSapPassword("fakeTokenPlaceholder")
                    .withSapPasswordUri("fakeTokenPlaceholder")
                    .withSapClientId("111")
                    .withSapPortNumber("1234")
                    .withSslPreference(SslPreference.ROOT_CERTIFICATE))
            .create();
    }
}
