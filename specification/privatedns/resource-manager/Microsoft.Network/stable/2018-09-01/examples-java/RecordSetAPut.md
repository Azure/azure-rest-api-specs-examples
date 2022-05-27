Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.privatedns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.privatedns.models.ARecord;
import com.azure.resourcemanager.privatedns.models.RecordType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for RecordSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/RecordSetAPut.json
     */
    /**
     * Sample code: PUT Private DNS Zone A Record Set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pUTPrivateDNSZoneARecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .createOrUpdateWithResponse(
                "resourceGroup1",
                "privatezone1.com",
                RecordType.A,
                "recordA",
                new RecordSetInner()
                    .withMetadata(mapOf("key1", "value1"))
                    .withTtl(3600L)
                    .withARecords(Arrays.asList(new ARecord().withIpv4Address("1.2.3.4"))),
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
