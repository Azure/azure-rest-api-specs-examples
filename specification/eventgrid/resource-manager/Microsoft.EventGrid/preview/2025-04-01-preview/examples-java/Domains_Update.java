
import com.azure.resourcemanager.eventgrid.models.Domain;
import com.azure.resourcemanager.eventgrid.models.InboundIpRule;
import com.azure.resourcemanager.eventgrid.models.IpActionType;
import com.azure.resourcemanager.eventgrid.models.PublicNetworkAccess;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Domains Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Domains_Update.
     * json
     */
    /**
     * Sample code: Domains_Update.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        Domain resource = manager.domains()
            .getByResourceGroupWithResponse("examplerg", "exampledomain1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withInboundIpRules(
                Arrays.asList(new InboundIpRule().withIpMask("12.18.30.15").withAction(IpActionType.ALLOW),
                    new InboundIpRule().withIpMask("12.18.176.1").withAction(IpActionType.ALLOW)))
            .apply();
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
