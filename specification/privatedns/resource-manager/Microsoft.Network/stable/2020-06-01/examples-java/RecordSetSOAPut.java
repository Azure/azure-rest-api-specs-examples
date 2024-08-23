
import com.azure.resourcemanager.privatedns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.privatedns.models.RecordType;
import com.azure.resourcemanager.privatedns.models.SoaRecord;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RecordSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetSOAPut.json
     */
    /**
     * Sample code: PUT Private DNS Zone SOA Record Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pUTPrivateDNSZoneSOARecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().createOrUpdateWithResponse("resourceGroup1",
            "privatezone1.com", RecordType.SOA, "@",
            new RecordSetInner().withMetadata(mapOf("key1", "fakeTokenPlaceholder")).withTtl(3600L)
                .withSoaRecord(new SoaRecord().withHost("azureprivatedns.net")
                    .withEmail("azureprivatedns-hostmaster.microsoft.com").withSerialNumber(1L).withRefreshTime(3600L)
                    .withRetryTime(300L).withExpireTime(2419200L)),
            null, null, com.azure.core.util.Context.NONE);
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
