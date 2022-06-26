import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualWanInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualWans CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualWANPut.json
     */
    /**
     * Sample code: VirtualWANCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualWANCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualWans()
            .createOrUpdate(
                "rg1",
                "wan1",
                new VirtualWanInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withDisableVpnEncryption(false)
                    .withTypePropertiesType("Basic"),
                Context.NONE);
    }

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
