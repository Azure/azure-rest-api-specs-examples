
import com.azure.resourcemanager.resourceconnector.models.AppliancePropertiesInfrastructureConfig;
import com.azure.resourcemanager.resourceconnector.models.Distro;
import com.azure.resourcemanager.resourceconnector.models.NetworkProfile;
import com.azure.resourcemanager.resourceconnector.models.Provider;
import com.azure.resourcemanager.resourceconnector.models.ProxyConfiguration;

/**
 * Samples for Appliances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesUpdateProxy.json
     */
    /**
     * Sample code: Update Appliance Proxy Configuration.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void updateApplianceProxyConfiguration(
        com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().define("appliance01").withRegion("West US").withExistingResourceGroup("testresourcegroup")
            .withDistro(Distro.AKSEDGE)
            .withInfrastructureConfig(new AppliancePropertiesInfrastructureConfig().withProvider(Provider.VMWARE))
            .withPublicKey("xxxxxxxx").withNetworkProfile(
                new NetworkProfile().withProxyConfiguration(new ProxyConfiguration().withVersion("latest")))
            .create();
    }
}
