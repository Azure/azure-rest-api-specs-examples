import com.azure.core.util.Context;
import com.azure.resourcemanager.dns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.dns.models.CnameRecord;
import com.azure.resourcemanager.dns.models.RecordType;
import java.util.HashMap;
import java.util.Map;

/** Samples for RecordSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateCNAMERecordset.json
     */
    /**
     * Sample code: Create CNAME recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createCNAMERecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .createOrUpdateWithResponse(
                "rg1",
                "zone1",
                "record1",
                RecordType.CNAME,
                new RecordSetInner()
                    .withMetadata(mapOf("key1", "value1"))
                    .withTtl(3600L)
                    .withCnameRecord(new CnameRecord().withCname("contoso.com")),
                null,
                null,
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
