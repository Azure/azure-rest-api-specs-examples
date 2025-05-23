
import com.azure.resourcemanager.hybridnetwork.models.OpenDeploymentResourceReference;
import com.azure.resourcemanager.hybridnetwork.models.ReferencedResource;
import com.azure.resourcemanager.hybridnetwork.models.SiteNetworkServicePropertiesFormat;
import com.azure.resourcemanager.hybridnetwork.models.Sku;
import com.azure.resourcemanager.hybridnetwork.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SiteNetworkServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * SiteNetworkServiceCreate.json
     */
    /**
     * Sample code: Create site network service.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createSiteNetworkService(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.siteNetworkServices().define("testSiteNetworkServiceName").withRegion("westUs2")
            .withExistingResourceGroup("rg1")
            .withProperties(new SiteNetworkServicePropertiesFormat()
                .withSiteReference(new ReferencedResource().withId(
                    "/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/sites/testSite"))
                .withNetworkServiceDesignVersionResourceReference(new OpenDeploymentResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkServiceDesignGroups/TestNetworkServiceDesignGroupName/networkServiceDesignVersions/1.0.0"))
                .withDesiredStateConfigurationGroupValueReferences(mapOf("MyVM_Configuration",
                    new ReferencedResource().withId(
                        "/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/configurationgroupvalues/MyVM_Configuration1"))))
            .withSku(new Sku().withName(SkuName.STANDARD)).create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
