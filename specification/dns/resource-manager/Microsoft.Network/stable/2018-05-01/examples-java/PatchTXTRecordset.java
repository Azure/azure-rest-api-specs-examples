
import com.azure.resourcemanager.dns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.dns.models.RecordType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RecordSets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchTXTRecordset.json
     */
    /**
     * Sample code: Patch TXT recordset.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchTXTRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().updateWithResponse("rg1", "zone1", "record1",
            RecordType.TXT, new RecordSetInner().withMetadata(mapOf("key2", "fakeTokenPlaceholder")), null,
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
