Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PrivateDnsZoneGroupInner;
import com.azure.resourcemanager.network.models.PrivateDnsZoneConfig;
import java.util.Arrays;

/** Samples for PrivateDnsZoneGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointDnsZoneGroupCreate.json
     */
    /**
     * Sample code: Create private dns zone group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPrivateDnsZoneGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateDnsZoneGroups()
            .createOrUpdate(
                "rg1",
                "testPe",
                "testPdnsgroup",
                new PrivateDnsZoneGroupInner()
                    .withPrivateDnsZoneConfigs(
                        Arrays
                            .asList(
                                new PrivateDnsZoneConfig()
                                    .withPrivateDnsZoneId(
                                        "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone1.com"))),
                Context.NONE);
    }
}
```
