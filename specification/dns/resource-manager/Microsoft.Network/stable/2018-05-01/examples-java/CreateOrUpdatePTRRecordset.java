import com.azure.resourcemanager.dns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.dns.models.PtrRecord;
import com.azure.resourcemanager.dns.models.RecordType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for RecordSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdatePTRRecordset.json
     */
    /**
     * Sample code: Create PTR recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPTRRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .createOrUpdateWithResponse(
                "rg1",
                "0.0.127.in-addr.arpa",
                "1",
                RecordType.PTR,
                new RecordSetInner()
                    .withMetadata(mapOf("key1", "fakeTokenPlaceholder"))
                    .withTtl(3600L)
                    .withPtrRecords(Arrays.asList(new PtrRecord().withPtrdname("localhost"))),
                null,
                null,
                com.azure.core.util.Context.NONE);
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
