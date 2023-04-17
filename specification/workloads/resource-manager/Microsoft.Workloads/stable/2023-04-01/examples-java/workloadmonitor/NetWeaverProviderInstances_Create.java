import com.azure.resourcemanager.workloads.models.SapNetWeaverProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;
import java.util.Arrays;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/NetWeaverProviderInstances_Create.json
     */
    /**
     * Sample code: Create a SAP monitor NetWeaver provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createASAPMonitorNetWeaverProvider(
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
                    .withSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename")
                    .withSslPreference(SslPreference.SERVER_CERTIFICATE))
            .create();
    }
}
