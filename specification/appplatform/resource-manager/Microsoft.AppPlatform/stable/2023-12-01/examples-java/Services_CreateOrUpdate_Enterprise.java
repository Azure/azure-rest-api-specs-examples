
import com.azure.resourcemanager.appplatform.fluent.models.ServiceResourceInner;
import com.azure.resourcemanager.appplatform.models.ClusterResourceProperties;
import com.azure.resourcemanager.appplatform.models.MarketplaceResource;
import com.azure.resourcemanager.appplatform.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Services_CreateOrUpdate_Enterprise.json
     */
    /**
     * Sample code: Services_CreateOrUpdate_Enterprise.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesCreateOrUpdateEnterprise(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().createOrUpdate("myResourceGroup", "myservice",
            new ServiceResourceInner().withLocation("eastus").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withProperties(new ClusterResourceProperties()
                    .withMarketplaceResource(new MarketplaceResource().withPlan("tanzu-asc-ent-mtr")
                        .withPublisher("vmware-inc").withProduct("azure-spring-cloud-vmware-tanzu-2")))
                .withSku(new Sku().withName("E0").withTier("Enterprise")),
            com.azure.core.util.Context.NONE);
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
