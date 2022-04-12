Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.dns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.dns.models.RecordType;
import com.azure.resourcemanager.dns.models.SoaRecord;
import java.util.HashMap;
import java.util.Map;

/** Samples for RecordSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateSOARecordset.json
     */
    /**
     * Sample code: Create SOA recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createSOARecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .createOrUpdateWithResponse(
                "rg1",
                "zone1",
                "@",
                RecordType.SOA,
                new RecordSetInner()
                    .withMetadata(mapOf("key1", "value1"))
                    .withTtl(3600L)
                    .withSoaRecord(
                        new SoaRecord()
                            .withHost("ns1.contoso.com")
                            .withEmail("hostmaster.contoso.com")
                            .withSerialNumber(1L)
                            .withRefreshTime(3600L)
                            .withRetryTime(300L)
                            .withExpireTime(2419200L)
                            .withMinimumTtl(300L)),
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
```
