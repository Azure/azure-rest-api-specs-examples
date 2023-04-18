import com.azure.resourcemanager.workloads.models.MsSqlServerProviderInstanceProperties;
import com.azure.resourcemanager.workloads.models.SslPreference;

/** Samples for ProviderInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/MsSqlServerProviderInstance_Create.json
     */
    /**
     * Sample code: Create a MsSqlServer provider.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createAMsSqlServerProvider(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .providerInstances()
            .define("myProviderInstance")
            .withExistingMonitor("myResourceGroup", "mySapMonitor")
            .withProviderSettings(
                new MsSqlServerProviderInstanceProperties()
                    .withHostname("hostname")
                    .withDbPort("5912")
                    .withDbUsername("user")
                    .withDbPassword("fakeTokenPlaceholder")
                    .withDbPasswordUri("fakeTokenPlaceholder")
                    .withSapSid("sid")
                    .withSslPreference(SslPreference.SERVER_CERTIFICATE)
                    .withSslCertificateUri("https://storageaccount.blob.core.windows.net/containername/filename"))
            .create();
    }
}
