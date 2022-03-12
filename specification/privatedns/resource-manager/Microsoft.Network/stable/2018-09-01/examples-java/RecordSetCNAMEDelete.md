Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.privatedns.models.RecordType;

/** Samples for RecordSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/RecordSetCNAMEDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone CNAME Record Set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZoneCNAMERecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .deleteWithResponse(
                "resourceGroup1", "privatezone1.com", RecordType.CNAME, "recordCNAME", null, Context.NONE);
    }
}
```
