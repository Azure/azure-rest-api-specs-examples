Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeletePTRRecordset.json
     */
    /**
     * Sample code: Delete PTR recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePTRRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .deleteWithResponse("rg1", "0.0.127.in-addr.arpa", "1", RecordType.PTR, null, Context.NONE);
    }
}
```
