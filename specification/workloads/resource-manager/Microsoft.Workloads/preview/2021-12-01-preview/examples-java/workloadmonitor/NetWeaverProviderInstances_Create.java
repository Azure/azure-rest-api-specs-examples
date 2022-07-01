import com.azure.resourcemanager.workloads.models.SapNetWeaverProviderInstanceProperties;
import java.util.Arrays;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/NetWeaverProviderInstances_Create.json
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
                    .withSapPassword("****")
                    .withSapPasswordUri("")
                    .withSapClientId("111")
                    .withSapPortNumber("1234")
                    .withSapSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename"))
            .create();
    }
}
