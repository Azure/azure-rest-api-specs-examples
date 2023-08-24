import com.azure.resourcemanager.resourceconnector.models.AppliancePropertiesInfrastructureConfig;
import com.azure.resourcemanager.resourceconnector.models.Distro;
import com.azure.resourcemanager.resourceconnector.models.Provider;

/** Samples for Appliances CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesCreate_Update.json
     */
    /**
     * Sample code: Create/Update Appliance.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void createUpdateAppliance(
        com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager
            .appliances()
            .define("appliance01")
            .withRegion("West US")
            .withExistingResourceGroup("testresourcegroup")
            .withDistro(Distro.AKSEDGE)
            .withInfrastructureConfig(new AppliancePropertiesInfrastructureConfig().withProvider(Provider.VMWARE))
            .create();
    }
}
