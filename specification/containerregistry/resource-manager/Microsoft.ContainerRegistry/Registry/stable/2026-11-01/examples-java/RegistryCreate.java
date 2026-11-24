
import com.azure.resourcemanager.containerregistry.fluent.models.RegistryInner;
import com.azure.resourcemanager.containerregistry.models.EndpointProtocol;
import com.azure.resourcemanager.containerregistry.models.RegionalEndpoints;
import com.azure.resourcemanager.containerregistry.models.Sku;
import com.azure.resourcemanager.containerregistry.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Registries Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/RegistryCreate.json
     */
    /**
     * Sample code: RegistryCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void registryCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().create("myResourceGroup", "myRegistry",
            new RegistryInner().withLocation("westus").withTags(mapOf("key", "fakeTokenPlaceholder"))
                .withSku(new Sku().withName(SkuName.PREMIUM)).withAdminUserEnabled(true).withDataEndpointEnabled(true)
                .withRegionalEndpoints(RegionalEndpoints.ENABLED).withEndpointProtocol(EndpointProtocol.IPV4AND_IPV6),
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
